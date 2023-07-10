package api

import (
	"github.com/deploifai/sdk-go/api/generated"
)

type API = GenericAPI[generated.GQLClient]

func NewAPI(gqlEndpoint string, restEndpoint string, headers RequestHeaders) API {

	return NewGenericAPI[generated.GQLClient](generated.NewClient, gqlEndpoint, restEndpoint, headers)
}

type Provider GenericProvider[generated.GQLClient]
