package azure

import (
	"context"
)

type IAMCredentials struct {
	AccessKey       string
	SecretAccessKey string
}

type Credentials struct {
	IAM IAMCredentials
}

type Client struct {
}

func CreateClient(ctx context.Context, creds Credentials, region string) (Client, error) {

	client := Client{}

	return client, nil
}
