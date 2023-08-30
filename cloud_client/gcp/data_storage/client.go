package data_storage

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/deploifai/sdk-go/cloud_client/gcp/root_provider"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
	"github.com/deploifai/sdk-go/cloud_client/utils"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io"
	"os"
	"path/filepath"
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

	file, err := os.Open(destAbsPath)
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

type pager struct {
	ctx          context.Context
	servicePager *storage.ObjectIterator
	maxResults   int
	done         bool
}

func (r *Client) NewListObjectsPager(input *implementable.ListObjectsInput) implementable.ListObjectsPager {

	var prefix = ""
	var startOffset = ""
	maxResults := utils.DataStorageListObjectsMaxResults

	if input != nil {
		if input.Prefix != nil {
			prefix = *input.Prefix
		}
		if input.Cursor != nil {
			startOffset = *input.Cursor
		}
		if input.MaxResultsPerPage != nil {
			maxResults = *input.MaxResultsPerPage
		}
	}

	query := &storage.Query{
		Prefix: prefix,
		//Delimiter:   utils.DataStorageObjectDelimiter,
		StartOffset: startOffset,
	}

	servicePager := r.service.Bucket(r.bucket).Objects(r.ctx, query)

	return &pager{ctx: r.ctx, servicePager: servicePager, maxResults: maxResults}
}

func (r *pager) NextPage(_ interface{}) (response implementable.ListObjectsResponse, err error) {

	for {
		object, err := r.servicePager.Next()
		if err == iterator.Done {
			r.done = true
			break
		}
		if err != nil {
			return implementable.ListObjectsResponse{}, err
		}

		if len(response.Objects) == r.maxResults {
			response.Cursor = &object.Name
			break
		}

		response.Objects = append(response.Objects, implementable.DataStorageObject{
			Key:  object.Name,
			Name: filepath.Base(object.Name),
		})
	}

	return response, nil
}

func (r *pager) More() bool {
	return !r.done
}
