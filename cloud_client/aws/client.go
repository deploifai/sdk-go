package aws

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

type Client struct {
	Config aws.Config
}

func CreateClient(ctx context.Context, creds Credentials, region string) (Client, error) {

	client := Client{}

	credentialsProvider := credentials.NewStaticCredentialsProvider(creds.IAM.AccessKey, creds.IAM.SecretAccessKey, "")

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(credentialsProvider),
		config.WithRegion(region),
	)
	if err != nil {
		return client, err
	}

	client.Config = cfg

	return client, nil
}
