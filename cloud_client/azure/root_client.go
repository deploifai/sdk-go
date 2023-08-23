package azure

import (
	"context"
)

type ServicePrincipalCredentials struct {
	SubscriptionId string
	TenantId       string
	ClientId       string
	ClientSecret   string
}

type Credentials struct {
	ServicePrincipal ServicePrincipalCredentials
}

type RootClient struct {
}

func NewRootClient(ctx context.Context, cred Credentials, region string) (RootClient, error) {

	client := RootClient{}

	return client, nil
}
