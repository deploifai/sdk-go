package data_storage

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/deploifai/sdk-go/cloud_client"
)

type DownloadFileInput struct {
	remoteObjectKey string
	destAbsPath     string
}

func (c *Client) DownloadFile(ctx context.Context, where generated.DataStorageContainerWhereUniqueInput, data DownloadFileInput) error {

	dataStorageContainerData, err := c.options.API.GetGQLClient().GetDataStorageContainer(ctx, where)
	if err != nil {
		return c.options.API.ProcessGQLError(err)
	}

	dataStorageContainerId := dataStorageContainerData.GetDataStorageContainer().GetID()
	dataStorageId := dataStorageContainerData.GetDataStorageContainer().GetDataStorageID()

	dataStorageData, err := c.options.API.GetGQLClient().GetDataStorage(ctx, generated.DataStorageWhereUniqueInput{ID: &dataStorageId})
	if err != nil {
		return c.options.API.ProcessGQLError(err)
	}

	cloudClientWrapper := cloud_client.NewCloudClientWrapper(ctx, c.options.API, *dataStorageData.GetDataStorage().GetCloudProfile().GetProvider())
	dataStorageClient, err := cloudClientWrapper.CloudClient.NewDataStorageClient(dataStorageId, dataStorageContainerId)
	if err != nil {
		return err
	}

	_, err = dataStorageClient.DownloadFile(data.remoteObjectKey, data.destAbsPath)

	return err
}
