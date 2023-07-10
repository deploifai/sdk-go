package config

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/credentials"
)

// LoadOptions are discrete set of options that are valid for loading the configuration
type LoadOptions struct {
	Credentials credentials.Provider
	API         api.Provider
}

// LoadOptionsFunc is a type alias for LoadOptions functional option
type LoadOptionsFunc func(*LoadOptions) error

func WithCredentials(v credentials.Provider) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.Credentials = v
		return nil
	}
}

func WithAPI(v api.Provider) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.API = v
		return nil
	}
}

func (o LoadOptions) getCredentials(ctx context.Context) (credentials.Provider, bool, error) {
	if o.Credentials == nil {
		return nil, false, nil
	}

	return o.Credentials, true, nil
}

func (o LoadOptions) getAPI(ctx context.Context) (api.Provider, bool, error) {
	if o.API == nil {
		return nil, false, nil
	}

	return o.API, true, nil
}
