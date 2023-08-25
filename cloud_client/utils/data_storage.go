package utils

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/api/generated"
)

const (
	DataStorageListObjectsMaxResults = 1000
)

func GetDataStorageAndContainer(ctx context.Context, api api.Provider, dataStorageId string, dataStorageContainerId string) (dataStorage generated.DataStorageFragment, dataStorageContainer generated.DataStorageContainerFragment, err error) {
	dataStorageData, err := api.GetGQLClient().GetDataStorage(ctx, generated.DataStorageWhereUniqueInput{ID: &dataStorageId})
	if err != nil {
		return dataStorage, dataStorageContainer, api.ProcessGQLError(err)
	}

	dataStorageContainerData, err := api.GetGQLClient().GetDataStorageContainer(ctx, generated.DataStorageContainerWhereUniqueInput{ID: &dataStorageContainerId})
	if err != nil {
		return dataStorage, dataStorageContainer, api.ProcessGQLError(err)
	}

	return *dataStorageData.GetDataStorage(), *dataStorageContainerData.GetDataStorageContainer(), nil
}
