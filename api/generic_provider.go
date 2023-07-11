package api

// GenericProvider is an interface for an api instance.
// The api instance should have a generated GQLClient and a default http RestClient.
// The api object should implement ProcessGQLError to process a GQLError
type GenericProvider[T GQLClient] interface {
	GetGQLClient() T
	GetRestClient() RestClient
	ProcessGQLError(err error) error
}
