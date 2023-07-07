package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Yamashou/gqlgenc/clientv2"
	"io"
	"net/http"
)

// GenericAPI is a generic API client that can be used to create a generated GQLClient and a default http RestClient
type GenericAPI[T GQLClient] struct {
	GQLClient T

	RestClient RestClient
}

// RequestHeader defines a request header for a http request
type RequestHeader struct {
	Key   string
	Value string
}

// RequestHeaders defines a list of RequestHeader
type RequestHeaders = []RequestHeader

// GQLClient defines the interface for the generated graphQL client
type GQLClient interface{}

// NewGQLClientFunc defines the generated function signature for creating a new GQLClient
type NewGQLClientFunc[T GQLClient] func(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) T

// createGQLClient takes a NewGQLClientFunc, endpoint and authToken
// runs the NewGQLClientFunc
// sets the authToken as a request header in the request interceptor
// returns a GQLClient
func createGQLClient[T GQLClient](newGQLClientFunc NewGQLClientFunc[T], endpoint string, headers RequestHeaders) T {

	requestInterceptor := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		for _, header := range headers {
			req.Header.Set(header.Key, header.Value)
		}

		return next(ctx, req, gqlInfo, res)
	}

	return newGQLClientFunc(http.DefaultClient, endpoint, nil, requestInterceptor)
}

// RestClient defines the interface for the default http client
type RestClient struct {
	endpoint   string
	headers    RequestHeaders
	httpClient *http.Client
}

// NewRequest takes a method, uri, headers and body
// if the request needs an empty body, like for a GET request, just pass in empty list of bytes
func (r *RestClient) NewRequest(method string, uri string, headers RequestHeaders, body []byte) (request *http.Request, err error) {
	if r.endpoint == "" {
		return nil, errors.New("rest client endpoint not set")
	}

	request, err = http.NewRequest(method, fmt.Sprintf("%s%s", r.endpoint, uri), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// set default headers
	for _, header := range r.headers {
		request.Header.Set(header.Key, header.Value)
	}

	// set custom headers
	for _, header := range headers {
		request.Header.Set(header.Key, header.Value)
	}

	return request, nil
}

func (r *RestClient) Do(request *http.Request) (*http.Response, error) {
	return r.httpClient.Do(request)
}

func (r *RestClient) ReadResponseJson(response *http.Response, t *interface{}) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
		// ignore error
	}(response.Body)

	return json.NewDecoder(response.Body).Decode(t)
}

func createRestClient(endpoint string, headers RequestHeaders) RestClient {
	return RestClient{
		endpoint:   endpoint,
		headers:    headers,
		httpClient: http.DefaultClient,
	}
}

// NewGenericAPI takes a NewGQLClientFunc, endpoint, and authToken, and
// returns an GenericAPI interface
// Example:
// ```
// api := NewGenericAPI[generated.GQLClient](generated.NewClient, "<endpoint>", "<authToken>")
// api.GetGQLClient().GetUser(...)
// ```
func NewGenericAPI[T GQLClient](newGQLClientFunc NewGQLClientFunc[T], gqlEndpoint string, restEndpoint string, headers RequestHeaders) GenericAPI[T] {

	return GenericAPI[T]{
		GQLClient:  createGQLClient[T](newGQLClientFunc, gqlEndpoint, headers),
		RestClient: createRestClient(restEndpoint, headers),
	}
}

func (api *GenericAPI[T]) GetGQLClient() T {
	return api.GQLClient
}

func (api *GenericAPI[T]) GetRestClient() RestClient {
	return api.RestClient
}

func (api *GenericAPI[T]) ProcessGQLError(err error) error {
	if handledError, ok := err.(*clientv2.ErrorResponse); ok {
		msg := "handled error: "
		if handledError.NetworkError != nil {
			msg = msg + fmt.Sprintf("network error: [status code = %d] %s\n", handledError.NetworkError.Code, handledError.NetworkError.Message)
		} else {
			msg = msg + fmt.Sprintf("graphql error: %v\n", handledError.GqlErrors)
		}
		return errors.New(msg)
	}

	return errors.New(fmt.Sprintf("unhandled error: %s\n", err.Error()))
}

// WithAuthHeader returns a RequestHeader with the key "Authorization" and the value of the authToken
func WithAuthHeader(authToken string) RequestHeader {
	return RequestHeader{Key: "Authorization", Value: authToken}
}
