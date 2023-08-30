package root_provider

type ServiceAccountCredentials struct {
	JsonKey []byte
}

type Credentials struct {
	ServiceAccount ServiceAccountCredentials
}

type RootProvider struct {
	Credentials Credentials
}

func New(cred Credentials) RootProvider {

	return RootProvider{Credentials: cred}
}
