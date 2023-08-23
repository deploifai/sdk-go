package cloud_client

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/deploifai/sdk-go/cloud_client/aws"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
)

type CloudClientWrapper struct {
	CloudClient implementable.CloudClient
	provider    generated.CloudProvider
}

func NewCloudClientWrapper(ctx context.Context, api api.Provider, provider generated.CloudProvider) (wrapper CloudClientWrapper) {
	switch provider {
	case generated.CloudProviderAws:
		client := aws.NewCloudClient(ctx, api)
		wrapper.CloudClient = &client
	case generated.CloudProviderAzure:
		//client := NewAzureCloudClient(ctx, api)
		//wrapper.CloudClient = &client
	case generated.CloudProviderGcp:
		//client := NewGCPCloudClient(ctx, api)
		//wrapper.CloudClient = &client
	}

	wrapper.provider = provider

	return wrapper
}
