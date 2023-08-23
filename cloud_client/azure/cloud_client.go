package azure

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/cloud_client/azure/data_storage"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
	"github.com/deploifai/sdk-go/cloud_client/utils"
)

type CloudClient struct {
	ctx context.Context
	api api.Provider
}

func NewCloudClient(ctx context.Context, api api.Provider) CloudClient {
	return CloudClient{
		ctx: ctx,
		api: api,
	}
}

func (r *CloudClient) NewDataStorageClient(dataStorageId string, dataStorageContainerId string) (dataStorageClient implementable.DataStorageClient, err error) {

	dataStorage, dataStorageContainer, err := utils.GetDataStorageAndContainer(r.ctx, r.api, dataStorageId, dataStorageContainerId)
	if err != nil {
		return dataStorageClient, err
	}

	container := dataStorageContainer.GetCloudName()
	azureConfig := dataStorage.GetCloudProviderYodaConfig().GetAzureConfig()

	return data_storage.New(r.ctx, *azureConfig.GetStorageAccount(), *azureConfig.GetStorageAccessKey(), *container)

}
