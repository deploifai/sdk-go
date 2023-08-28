package dataset

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

type UploadFileInput struct {
	srcAbspath      string
	remoteObjectKey string
}

func (c *Client) UploadFile(ctx context.Context, where generated.DataStorageContainerWhereUniqueInput, data UploadFileInput) error {

	dataStorage, dataStorageContainer, err := getDataStorageAndContainer(ctx, c.options.API, where)
	if err != nil {
		return err
	}

	dataStorageClient, err := newDataStorageClient(ctx, c.options.API, dataStorage, dataStorageContainer)
	if err != nil {
		return err
	}

	_, err = dataStorageClient.UploadFile(data.srcAbspath, data.remoteObjectKey)

	return err
}
