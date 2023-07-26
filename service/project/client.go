package project

import (
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/config"
	"github.com/deploifai/sdk-go/credentials"
)

type Client struct {
	options Options
}

type Options struct {
	Credentials credentials.Provider
	API         api.Provider
}

func New(options Options) *Client {
	return &Client{options: options}
}

func NewFromConfig(cfg config.Config) *Client {
	opts := Options{
		Credentials: cfg.Credentials,
		API:         cfg.API,
	}

	return New(opts)
}
