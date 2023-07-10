package config

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/api/host"
	"github.com/deploifai/sdk-go/credentials"
)

// An configResolver will extract configuration data from the configs slice
// using the provider interfaces to extract specific functionality. The extracted
// configuration values will be written to the Config value.
//
// The resolver should return an error if it fails to extract the data, the
// data is malformed, or incomplete.
type configResolver func(ctx context.Context, cfg *Config, configs configs) error

var defaultConfigResolvers = []configResolver{
	resolveCredentials,
	resolveAPI,
}

func resolveCredentials(ctx context.Context, cfg *Config, configs configs) error {
	c, found, err := getCredentials(ctx, configs)
	if err != nil {
		return err
	}
	if !found {
		cfg.Credentials = credentials.NewCredentials("")
	}

	cfg.Credentials = c
	return nil
}

func resolveAPI(ctx context.Context, cfg *Config, configs configs) error {
	c, found, err := getAPI(ctx, configs)
	if err != nil {
		return err
	}
	if !found {
		creds, err := cfg.Credentials.Retrieve()
		if err != nil {
			return err
		}

		headers := api.RequestHeaders{}
		if creds.AuthToken != "" {
			headers = append(headers, api.WithAuthHeader(creds.AuthToken))
		}

		cfg.API = api.NewAPI(host.Endpoint.GraphQL, host.Endpoint.Rest.Base, headers)
		return nil
	}

	cfg.API = c
	return nil
}
