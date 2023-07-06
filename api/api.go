package api

import (
	"context"
	"github.com/Yamashou/gqlgenc/clientv2"
	"net/http"
)

type API[T GQLClient] struct {
	Endpoint string

	AuthToken string

	Client T
}

// GQLClient defines the interface for the generated client
type GQLClient interface{}

// NewGQLClientFunc defines the generated function signature for creating a new GQLClient
type NewGQLClientFunc[T GQLClient] func(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) T

// createGQLClient takes a NewGQLClientFunc, endpoint and authToken
// runs the NewGQLClientFunc
// sets the authToken as a request header in the request interceptor
// returns a GQLClient
func createGQLClient[T GQLClient](newGQLClientFunc NewGQLClientFunc[T], endpoint string, authToken string) T {

	requestInterceptor := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		req.Header.Set("Authorization", authToken)

		return next(ctx, req, gqlInfo, res)
	}

	return newGQLClientFunc(http.DefaultClient, endpoint, nil, requestInterceptor)
}

// NewAPI takes a NewGQLClientFunc, endpoint, and authToken, and
// returns an API interface
// Example:
// ```
// api := NewAPI[generated.GQLClient](generated.NewClient, "<endpoint>", "<authToken>")
// api.GetClient().GetUser(...)
// ```
func NewAPI[T GQLClient](newGQLClientFunc NewGQLClientFunc[T], endpoint string, authToken string) API[T] {

	return API[T]{
		Endpoint:  endpoint,
		AuthToken: authToken,
		Client:    createGQLClient[T](newGQLClientFunc, endpoint, authToken),
	}
}
