package aws

import (
	"context"
	"github.com/deploifai/sdk-go/api"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/deploifai/sdk-go/cloud_client/aws/data_storage"
	"github.com/deploifai/sdk-go/cloud_client/aws/root_client"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
)

type CloudClient struct {
	ctx context.Context
	api api.Provider
}

func NewCloudClient(ctx context.Context, api api.Provider) CloudClient {
	return CloudClient{
		ctx: ctx,
		api: api,
	}
}

func (r *CloudClient) NewDataStorageClient(dataStorageId string, dataStorageContainerId string) (dataStorageClient implementable.DataStorageClient, err error) {

	dataStorageData, err := r.api.GetGQLClient().GetDataStorage(r.ctx, generated.DataStorageWhereUniqueInput{ID: &dataStorageId})
	if err != nil {
		return dataStorageClient, r.api.ProcessGQLError(err)
	}

	dataStorageContainerData, err := r.api.GetGQLClient().GetDataStorageContainer(r.ctx, generated.DataStorageContainerWhereUniqueInput{ID: &dataStorageContainerId})
	if err != nil {
		return dataStorageClient, r.api.ProcessGQLError(err)
	}

	dataStorage := dataStorageData.GetDataStorage()
	dataStorageContainer := dataStorageContainerData.GetDataStorageContainer()

	bucket := dataStorageContainer.GetCloudName()
	awsConfig := dataStorage.GetCloudProviderYodaConfig().GetAwsConfig()

	rootClient, err := root_client.NewRootClient(
		r.ctx,
		root_client.Credentials{
			IAM: root_client.IAMCredentials{
				AccessKey:       *awsConfig.GetAwsAccessKey(),
				SecretAccessKey: *awsConfig.GetAwsSecretAccessKey()},
		},
		awsConfig.GetAwsRegion())
	if err != nil {
		return dataStorageClient, err
	}

	return data_storage.New(r.ctx, &rootClient, *bucket), nil
}
