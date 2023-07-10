package api

// GenericProvider is an interface for an api object.
// The api object is expected to have a generated GQLClient and a default http RestClient.
// The api object should implement a method to process a GQLError
type GenericProvider[T GQLClient] interface {
	GetGQLClient() T
	GetRestClient() RestClient
	ProcessGQLError(err error) error
}
