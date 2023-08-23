package data_storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/deploifai/sdk-go/cloud_client/aws/root_provider"
	"io"
	"os"
)

// Client is a wrapper around the AWS S3 client.
// It implements the implementable.DataStorageClient interface.
type Client struct {
	ctx     context.Context
	service s3.Client
	bucket  string
}

func New(ctx context.Context, rootProvider *root_provider.RootProvider, bucket string) *Client {

	service := s3.NewFromConfig(rootProvider.Config)

	return &Client{ctx: ctx, service: *service, bucket: bucket}
}

// UploadFile uploads a file to the S3 bucket given a source absolute path and a remote object key.
func (r *Client) UploadFile(srcAbsPath string, remoteObjectKey string) (interface{}, error) {

	file, err := os.Open(srcAbsPath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	object, err := r.service.PutObject(r.ctx, &s3.PutObjectInput{
		Bucket: &r.bucket,
		Key:    &remoteObjectKey,
		Body:   file,
	})
	if err != nil {
		return nil, err
	}

	return *object, nil
}

// DownloadFile downloads a file from the S3 bucket given a remote object key and a destination absolute path.
func (r *Client) DownloadFile(remoteObjectKey string, destAbsPath string) (interface{}, error) {

	object, err := r.service.GetObject(r.ctx, &s3.GetObjectInput{
		Bucket: &r.bucket,
		Key:    &remoteObjectKey,
	})
	if err != nil {
		return nil, err
	}

	file, err := os.Create(destAbsPath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, object.Body)
	if err != nil {
		return nil, err
	}

	return *object, nil
}
