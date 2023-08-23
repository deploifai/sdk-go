package root_client

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type IAMCredentials struct {
	AccessKey       string
	SecretAccessKey string
}

type Credentials struct {
	IAM IAMCredentials
}

type RootClient struct {
	Config aws.Config
}

func NewRootClient(ctx context.Context, cred Credentials, region string) (RootClient, error) {

	rootClient := RootClient{}

	credentialsProvider := credentials.NewStaticCredentialsProvider(cred.IAM.AccessKey, cred.IAM.SecretAccessKey, "")

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(credentialsProvider),
		config.WithRegion(region),
	)
	if err != nil {
		return rootClient, err
	}

	rootClient.Config = cfg

	return rootClient, nil
}
