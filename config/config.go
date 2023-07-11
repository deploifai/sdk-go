package config

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/credentials"
)

// Config provides the configuration for the service clients.
type Config struct {
	Credentials credentials.Provider
	API         api.Provider
}

// AppendFromLoaders iterates over the slice of loaders passed in calling each
// loader function in order. The external config value returned by the loader
// will be added to the returned configs slice.
//
// If a loader returns an error this method will stop iterating and return
// that error.
func (cs configs) AppendFromLoaders(ctx context.Context, loaders []loader) (configs, error) {
	for _, fn := range loaders {
		cfg, err := fn(ctx, cs)
		if err != nil {
			return nil, err
		}

		cs = append(cs, cfg)
	}

	return cs, nil
}

// ResolveConfig returns a configuration populated with values by calling
// the resolvers slice passed in. Each resolver is called in order. Any resolver
// may overwrite the configuration value of a previous resolver.
//
// If a resolver returns an error this method will return that error, and stop
// iterating over the resolvers.
func (cs configs) ResolveConfig(ctx context.Context, resolvers []configResolver) (Config, error) {
	var cfg Config

	for _, fn := range resolvers {
		if err := fn(ctx, &cfg, cs); err != nil {
			return Config{}, err
		}
	}

	return cfg, nil
}

// LoadDefaultConfig loads the default configuration for the service clients.
func LoadDefaultConfig(ctx context.Context, optFns ...func(*LoadOptions) error) (Config, error) {

	var options LoadOptions

	for _, fn := range optFns {
		if err := fn(&options); err != nil {
			return Config{}, err
		}
	}

	// assign LoadOptions to configs
	var cfgCpy = configs{options}

	cfgCpy, err := cfgCpy.AppendFromLoaders(ctx, defaultLoaders)
	if err != nil {
		return Config{}, err
	}

	cfg, err := cfgCpy.ResolveConfig(ctx, defaultConfigResolvers)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
