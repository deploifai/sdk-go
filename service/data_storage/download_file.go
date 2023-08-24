package data_storage

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

type DownloadFileInput struct {
	remoteObjectKey string
	destAbsPath     string
}

func (c *Client) DownloadFile(ctx context.Context, where generated.DataStorageContainerWhereUniqueInput, data DownloadFileInput) error {

	dataStorage, dataStorageContainer, err := getDataStorageAndContainer(ctx, c.options.API, where)
	if err != nil {
		return err
	}

	dataStorageClient, err := newDataStorageClient(ctx, c.options.API, dataStorage, dataStorageContainer)
	if err != nil {
		return err
	}

	_, err = dataStorageClient.DownloadFile(data.remoteObjectKey, data.destAbsPath)

	return err
}
