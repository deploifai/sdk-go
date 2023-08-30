package dataset

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
	"os"
)

type DownloadFileInput struct {
	RemoteObjectKey string
	DestAbsPath     string
}

func (c *Client) DownloadFile(ctx context.Context, where generated.DataStorageWhereUniqueInput, data DownloadFileInput) error {

	dataStorage, dataStorageContainer, err := getDataStorageAndContainer(ctx, c.options.API, where)
	if err != nil {
		return err
	}

	dataStorageClient, err := newDataStorageClient(ctx, c.options.API, dataStorage, dataStorageContainer)
	if err != nil {
		return err
	}

	if _, err := os.Create(data.DestAbsPath); err != nil {
		return err
	}

	_, err = dataStorageClient.DownloadFile(data.RemoteObjectKey, data.DestAbsPath)

	return err
}
