package credentials

import "context"

// Credentials holds the credentials needed to authenticate API requests.
type Credentials struct {
	AuthToken string
}

// Provider is the interface for any component which will provide
// credentials.
type Provider interface {
	Retrieve() (Credentials, error)
}

// ProviderFunc provides a helper wrapping a function value to
// satisfy the Provider interface.
type ProviderFunc func(context.Context) (Credentials, error)

// Retrieve delegates to the function value the ProviderFunc wraps.
func (fn ProviderFunc) Retrieve(ctx context.Context) (Credentials, error) {
	return fn(ctx)
}

// NewCredentials returns a new Credentials instance.
func NewCredentials(authToken string) Credentials {
	return Credentials{
		AuthToken: authToken,
	}
}

func (c Credentials) Retrieve() (Credentials, error) {
	return c, nil
}
