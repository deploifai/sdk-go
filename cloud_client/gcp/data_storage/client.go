package data_storage

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/deploifai/sdk-go/cloud_client/gcp/root_provider"
	"google.golang.org/api/option"
	"io"
	"os"
)

type Client struct {
	ctx     context.Context
	service storage.Client
	bucket  string
}

func New(ctx context.Context, rootProvider root_provider.RootProvider, bucket string) (client *Client, err error) {

	option := option.WithCredentialsJSON(rootProvider.Credentials.ServiceAccount.JsonKey)
	service, err := storage.NewClient(ctx, option)
	if err != nil {
		return nil, err
	}

	return &Client{ctx: ctx, service: *service, bucket: bucket}, nil

}

func (r *Client) UploadFile(srcAbsPath string, remoteObjectKey string) (interface{}, error) {

	file, err := os.Open(srcAbsPath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	wc := r.service.Bucket(r.bucket).Object(remoteObjectKey).NewWriter(r.ctx)
	result, err := io.Copy(wc, file)
	if err != nil {
		return nil, err
	}
	if err := wc.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Client) DownloadFile(remoteObjectKey string, destAbsPath string) (interface{}, error) {

	rc, err := r.service.Bucket(r.bucket).Object(remoteObjectKey).NewReader(r.ctx)
	if err != nil {
		return nil, err
	}
	defer func(rc *storage.Reader) {
		_ = rc.Close()
	}(rc)

	file, err := os.Create(destAbsPath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	result, err := io.Copy(file, rc)
	if err != nil {
		return nil, err
	}

	return result, nil
}
