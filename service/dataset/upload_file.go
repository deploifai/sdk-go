package dataset

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

type UploadFileInput struct {
	SrcAbspath      string
	RemoteObjectKey string
}

func (c *Client) UploadFile(ctx context.Context, where generated.DataStorageWhereUniqueInput, data UploadFileInput) error {

	dataStorage, dataStorageContainer, err := getDataStorageAndContainer(ctx, c.options.API, where)
	if err != nil {
		return err
	}

	dataStorageClient, err := newDataStorageClient(ctx, c.options.API, dataStorage, dataStorageContainer)
	if err != nil {
		return err
	}

	_, err = dataStorageClient.UploadFile(data.SrcAbspath, data.RemoteObjectKey)

	return err
}
