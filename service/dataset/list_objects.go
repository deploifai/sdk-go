package dataset

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
)

type ListObjectsInput = implementable.ListObjectsInput

func (c *Client) NewListObjectsPager(ctx context.Context, where generated.DataStorageWhereUniqueInput, input *ListObjectsInput) (pager implementable.ListObjectsPager, err error) {

	dataStorage, dataStorageContainer, err := getDataStorageAndContainer(ctx, c.options.API, where)
	if err != nil {
		return pager, err
	}

	dataStorageClient, err := newDataStorageClient(ctx, c.options.API, dataStorage, dataStorageContainer)
	if err != nil {
		return pager, err
	}

	return dataStorageClient.NewListObjectsPager(input), nil

}
