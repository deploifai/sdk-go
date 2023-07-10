# SDK-GO

This is the Go SDK for Deploifai. It is used to interact with the Deploifai API.

## Installation

```shell
go get github.com/deploifai/sdk-go
```

## Usage

```go
package main

import (
	"context"
	"github.com/deploifai/sdk-go/config"
	"github.com/deploifai/sdk-go/credentials"
	"github.com/deploifai/sdk-go/service/workspace"
)

func main() {
	// Create a new Credentials instance
	creds := credentials.NewCredentials("<deploifai auth token>")

	// Create a new config with the default config loader, and use the credentials
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithCredentials(creds))

	// Create a new workspace client
	client := workspace.NewFromConfig(cfg)

	// List all workspaces
	_, _ = client.List(context.TODO())
}
```
