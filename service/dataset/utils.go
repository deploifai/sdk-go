package dataset

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/deploifai/sdk-go/cloud_client"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
	"path/filepath"
)

func getDataStorageAndContainer(ctx context.Context, api api.Provider, where generated.DataStorageWhereUniqueInput) (dataStorage generated.DataStorageFragment, dataStorageContainer generated.DataStorageContainerFragment, err error) {
	data, err := api.GetGQLClient().GetDataStorage(ctx, where)
	if err != nil {
		return dataStorage, dataStorageContainer, api.ProcessGQLError(err)
	}

	dataStorage = *data.GetDataStorage()

	return dataStorage, *dataStorage.GetContainers()[0], nil
}

func newDataStorageClient(ctx context.Context, api api.Provider, dataStorage generated.DataStorageFragment, dataStorageContainer generated.DataStorageContainerFragment) (implementable.DataStorageClient, error) {
	cloudClientWrapper := cloud_client.NewCloudClientWrapper(ctx, api, *dataStorage.GetCloudProfile().GetProvider())
	return cloudClientWrapper.CloudClient.NewDataStorageClient(dataStorage.GetID(), dataStorageContainer.GetID())
}

func cleanRemoteObjectPrefix(raw string) string {
	f := filepath.ToSlash(filepath.Clean(raw))
	// remove leading slash if any
	if len(f) > 0 && f[0] == '/' {
		f = f[1:]
	}
	// if at root, then it should just be an empty string
	// if not at root, then it should end with a slash
	if f == "." || f == "" {
		f = ""
	} else {
		f = f + "/"
	}
	return f
}

type workerTask struct {
	path       string
	resultChan chan<- interface{}
	errChan    chan<- error
}
