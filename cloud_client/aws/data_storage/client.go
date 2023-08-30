package data_storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/deploifai/sdk-go/cloud_client/aws/root_provider"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
	"github.com/deploifai/sdk-go/cloud_client/utils"
	"io"
	"os"
	"path/filepath"
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

	// create the file
	file, err := os.Create(destAbsPath)
	if err != nil {
		return nil, err
	}

	// close the file after it is no longer required
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, object.Body)
	if err != nil {
		return nil, err
	}

	return *object, nil
}

type pager struct {
	ctx          context.Context
	servicePager *s3.ListObjectsV2Paginator
}

func (r *Client) NewListObjectsPager(input *implementable.ListObjectsInput) implementable.ListObjectsPager {

	var prefix *string
	var maxKeys int32 = utils.DataStorageListObjectsMaxResults
	var startAfter *string

	if input != nil {
		prefix = input.Prefix
		if input.MaxResultsPerPage != nil {
			maxKeys = int32(*input.MaxResultsPerPage)
		}
		if input.Cursor != nil {
			startAfter = input.Cursor
		}
	}

	params := &s3.ListObjectsV2Input{
		Bucket:     &r.bucket,
		Prefix:     prefix,
		MaxKeys:    maxKeys,
		StartAfter: startAfter,
	}

	servicePager := s3.NewListObjectsV2Paginator(&r.service, params)

	return &pager{ctx: r.ctx, servicePager: servicePager}
}

func (r *pager) NextPage(_ interface{}) (response implementable.ListObjectsResponse, err error) {
	resp, err := r.servicePager.NextPage(r.ctx)
	if err != nil {
		return response, err
	}

	for _, v := range resp.Contents {
		response.Objects = append(response.Objects, implementable.DataStorageObject{
			Key:  *v.Key,
			Name: filepath.Base(*v.Key),
		})
	}

	if resp.IsTruncated {
		response.Cursor = resp.Contents[len(resp.Contents)-1].Key
	}

	return response, nil

}

func (r *pager) More() bool {
	return r.servicePager.HasMorePages()
}
