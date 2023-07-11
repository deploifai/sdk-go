package config

import (
	"context"
	"github.com/deploifai/sdk-go/credentials"
)

// IConfig is an interface that all external configuration values must satisfy.
type IConfig interface{}

// configs is a slice of IConfig types.
type configs []IConfig

type loader func(context.Context, configs) (IConfig, error)

// defaultLoaders are a slice of functions that will read external configuration
// sources for configuration values. These values are read by the configResolver(s)
// using interfaces to extract specific information from the external configuration.
var defaultLoaders = []loader{
	loadEnvConfig,
}

type EnvConfig struct {
	Credentials credentials.Credentials
}

// loadEnvConfig reads configuration values from the OS's environment variables.
// Returning a IConfig typed EnvConfig to satisfy the loader func type.
func loadEnvConfig(ctx context.Context, cfgs configs) (IConfig, error) {
	return NewEnvConfig()
}

func NewEnvConfig() (EnvConfig, error) {
	return EnvConfig{}, nil
}
