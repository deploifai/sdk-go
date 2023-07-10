package credentials

import "context"

type Credentials struct {
	AuthToken string
}

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

func NewCredentials(authToken string) Credentials {
	return Credentials{
		AuthToken: authToken,
	}
}

func (c Credentials) Retrieve() (Credentials, error) {
	return c, nil
}
