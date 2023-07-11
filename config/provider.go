package config

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/credentials"
)

// CredentialsProvider is an interface for retrieving an credentials.Provider
type CredentialsProvider interface {
	getCredentials(ctx context.Context) (credentials.Provider, bool, error)
}

func getCredentials(ctx context.Context, configs configs) (v credentials.Provider, found bool, err error) {
	for _, config := range configs {
		if p, ok := config.(CredentialsProvider); ok {
			v, found, err = p.getCredentials(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}

// APIProvider is an interface for retrieving an api.Provider
type APIProvider interface {
	getAPI(ctx context.Context) (api.Provider, bool, error)
}

// getAPI searches the slice of configs and returns the API set on configs
func getAPI(ctx context.Context, configs configs) (v api.Provider, found bool, err error) {
	for _, config := range configs {
		if p, ok := config.(APIProvider); ok {
			v, found, err = p.getAPI(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}
