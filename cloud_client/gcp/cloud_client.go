package gcp

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/cloud_client/gcp/data_storage"
	"github.com/deploifai/sdk-go/cloud_client/gcp/root_provider"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
	"github.com/deploifai/sdk-go/cloud_client/utils"
)

type CloudClient struct {
	ctx context.Context
	api api.Provider
}

func NewCloudClient(ctx context.Context, api api.Provider) CloudClient {
	return CloudClient{ctx: ctx, api: api}
}

func (r *CloudClient) NewDataStorageClient(dataStorageId string, dataStorageContainerId string) (dataStorageClient implementable.DataStorageClient, err error) {

	dataStorage, dataStorageContainer, err := utils.GetDataStorageAndContainer(r.ctx, r.api, dataStorageId, dataStorageContainerId)
	if err != nil {
		return dataStorageClient, err
	}

	bucket := dataStorageContainer.GetCloudName()
	gcpConfig := dataStorage.GetCloudProviderYodaConfig().GetGcpConfig()

	jsonKey := []byte(*gcpConfig.GetGcpServiceAccountKey())
	rootProvider := root_provider.New(root_provider.Credentials{ServiceAccount: root_provider.ServiceAccountCredentials{JsonKey: jsonKey}})

	return data_storage.New(r.ctx, rootProvider, *bucket)
}
