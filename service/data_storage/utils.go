package data_storage

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/deploifai/sdk-go/cloud_client"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
)

func getDataStorageAndContainer(ctx context.Context, api api.Provider, where generated.DataStorageContainerWhereUniqueInput) (dataStorage generated.DataStorageFragment, dataStorageContainer generated.DataStorageContainerFragment, err error) {
	dataStorageContainerData, err := api.GetGQLClient().GetDataStorageContainer(ctx, where)
	if err != nil {
		return dataStorage, dataStorageContainer, api.ProcessGQLError(err)
	}

	dataStorageId := dataStorageContainerData.GetDataStorageContainer().GetDataStorageID()

	dataStorageData, err := api.GetGQLClient().GetDataStorage(ctx, generated.DataStorageWhereUniqueInput{ID: &dataStorageId})
	if err != nil {
		return dataStorage, dataStorageContainer, api.ProcessGQLError(err)
	}

	return *dataStorageData.GetDataStorage(), *dataStorageContainerData.GetDataStorageContainer(), nil
}

func newDataStorageClient(ctx context.Context, api api.Provider, dataStorage generated.DataStorageFragment, dataStorageContainer generated.DataStorageContainerFragment) (implementable.DataStorageClient, error) {
	cloudClientWrapper := cloud_client.NewCloudClientWrapper(ctx, api, *dataStorage.GetCloudProfile().GetProvider())
	return cloudClientWrapper.CloudClient.NewDataStorageClient(dataStorage.GetID(), dataStorageContainer.GetID())
}
