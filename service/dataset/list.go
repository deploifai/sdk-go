package dataset

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

func (c *Client) List(ctx context.Context, whereAccount generated.AccountWhereUniqueInput, whereDataStorage *generated.DataStorageWhereInput) (dataStorages []generated.DataStorageFragment, err error) {

	data, err := c.options.API.GetGQLClient().GetDataStorages(ctx, whereAccount, whereDataStorage)
	if err != nil {
		return nil, c.options.API.ProcessGQLError(err)
	}
	for _, dataStorage := range data.GetDataStorages() {
		dataStorages = append(dataStorages, *dataStorage)
	}

	return dataStorages, nil
}
