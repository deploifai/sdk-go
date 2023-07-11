package api

import (
	"github.com/deploifai/sdk-go/api/generated"
)

// API is a type alias for GenericAPI[generated.GQLClient].
// The generated.GQLClient is generated based on graphQL query schema in this module.
type API = GenericAPI[generated.GQLClient]

// NewAPI calls NewGenericAPI[generated.GQLClient] function with generated.NewClient as the first argument.
func NewAPI(gqlEndpoint string, restEndpoint string, headers RequestHeaders) API {

	return NewGenericAPI[generated.GQLClient](generated.NewClient, gqlEndpoint, restEndpoint, headers)
}

// Provider is a type alias for GenericProvider[generated.GQLClient].
type Provider GenericProvider[generated.GQLClient]
