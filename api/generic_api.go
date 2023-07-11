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

// GenericAPI is a generic api client that uses a generated GQLClient and a default http RestClient.
type GenericAPI[T GQLClient] struct {
	GQLClient T

	RestClient RestClient
}

// RequestHeader defines a request header for a http request.
type RequestHeader struct {
	Key   string
	Value string
}

// RequestHeaders defines a list of RequestHeader(s).
type RequestHeaders = []RequestHeader

// GQLClient defines the interface for the generated graphQL client.
type GQLClient interface{}

// NewGQLClientFunc defines the signature of the generated function that creates a new GQLClient.
// The function is generated from the [gqlgenc] tool.
// [gqlgenc]: https://github.com/Yamashou/gqlgenc
type NewGQLClientFunc[T GQLClient] func(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) T

// createGQLClient takes a NewGQLClientFunc, endpoint and request headers, and
// runs the NewGQLClientFunc, passing in the endpoint.
// It also sets the request headers in the request interceptor.
// And returns a GQLClient.
func createGQLClient[T GQLClient](newGQLClientFunc NewGQLClientFunc[T], endpoint string, headers RequestHeaders) T {

	requestInterceptor := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		for _, header := range headers {
			req.Header.Set(header.Key, header.Value)
		}

		return next(ctx, req, gqlInfo, res)
	}

	return newGQLClientFunc(http.DefaultClient, endpoint, nil, requestInterceptor)
}

// RestClient defines the interface for the default http client.
type RestClient struct {
	endpoint   string
	headers    RequestHeaders
	httpClient *http.Client
}

// NewRequest takes a method, uri, request headers and body.
// if the request needs an empty body, like for a GET request, just pass in an empty list of bytes.
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

// Do sends a http request and returns a http response.
func (r *RestClient) Do(request *http.Request) (*http.Response, error) {
	return r.httpClient.Do(request)
}

// ReadResponseJson reads the response body and unmarshals it into the given interface.
func (r *RestClient) ReadResponseJson(response *http.Response, v any) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
		// ignore error
	}(response.Body)

	return json.NewDecoder(response.Body).Decode(v)
}

func createRestClient(endpoint string, headers RequestHeaders) RestClient {
	return RestClient{
		endpoint:   endpoint,
		headers:    headers,
		httpClient: http.DefaultClient,
	}
}

// NewGenericAPI takes a NewGQLClientFunc, a gqlEndpoint for the GQLClient, a restEndpoint for the RestClient, and request headers.
// The request headers are set the same for both clients. NewGenericAPI returns an GenericAPI interface.
// Example:
//
//	api := NewGenericAPI[generated.GQLClient](generated.NewClient, "<gqlEndpoint>", "<restEndpoint>", RequestHeaders{WithAuthToken("<authToken>")})
func NewGenericAPI[T GQLClient](newGQLClientFunc NewGQLClientFunc[T], gqlEndpoint string, restEndpoint string, headers RequestHeaders) GenericAPI[T] {

	return GenericAPI[T]{
		GQLClient:  createGQLClient[T](newGQLClientFunc, gqlEndpoint, headers),
		RestClient: createRestClient(restEndpoint, headers),
	}
}

// GetGQLClient returns the GQLClient.
func (r GenericAPI[T]) GetGQLClient() T {
	return r.GQLClient
}

// GetRestClient returns the RestClient.
func (r GenericAPI[T]) GetRestClient() RestClient {
	return r.RestClient
}

// ProcessGQLError takes an error returned by the GQLClient, improve it's message formatting and returns it.
func (r GenericAPI[T]) ProcessGQLError(err error) error {
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
