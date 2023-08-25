package data_storage

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
	"github.com/deploifai/sdk-go/cloud_client/utils"
	"os"
	"path/filepath"
)

// Client is a wrapper around the Azure Blob Storage client.
// It implements the implementable.DataStorageClient interface.
type Client struct {
	ctx       context.Context
	service   azblob.Client
	container string
}

func New(ctx context.Context, accountName string, accountKey string, container string) (client *Client, err error) {

	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return client, err
	}

	service, err := azblob.NewClientWithSharedKeyCredential(serviceURL, cred, nil)
	if err != nil {
		return client, err
	}

	return &Client{ctx: ctx, service: *service, container: container}, nil
}

// UploadFile uploads a file to the Azure Blob Storage container given a source absolute path and a remote object key.
func (r *Client) UploadFile(srcAbsPath string, remoteObjectKey string) (interface{}, error) {

	// open the file
	file, err := os.Open(srcAbsPath)
	if err != nil {
		return nil, err
	}

	// close the file after it is no longer required
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// upload the file
	response, err := r.service.UploadFile(r.ctx, r.container, remoteObjectKey, file,
		&azblob.UploadFileOptions{
			BlockSize:   int64(1024),
			Concurrency: uint16(0),
		})
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DownloadFile downloads a file from the Azure Blob Storage container given a remote object key and a destination absolute path.
func (r *Client) DownloadFile(remoteObjectKey string, destAbsPath string) (interface{}, error) {

	// create the file
	file, err := os.Create(destAbsPath)
	if err != nil {
		return nil, err
	}

	// close the file after it is no longer required
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// download the file
	response, err := r.service.DownloadFile(r.ctx, r.container, remoteObjectKey, file, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type pager struct {
	ctx          context.Context
	servicePager *runtime.Pager[azblob.ListBlobsFlatResponse]
}

func (r *Client) NewListObjectsPager(input *implementable.ListObjectsInput) implementable.ListObjectsPager {

	var marker *string = nil
	var maxResults int32 = utils.DataStorageListObjectsMaxResults
	var prefix *string = nil

	if input != nil {
		marker = input.Cursor
		if input.MaxResultsPerPage != nil {
			maxResults = int32(*input.MaxResultsPerPage)
		}
		prefix = input.Prefix
	}

	options := azblob.ListBlobsFlatOptions{
		Marker:     marker,
		MaxResults: &maxResults,
		Prefix:     prefix,
	}

	servicePager := r.service.NewListBlobsFlatPager(r.container, &options)

	return &pager{ctx: r.ctx, servicePager: servicePager}

}

func (r *pager) NextPage(_ interface{}) (response implementable.ListObjectsResponse, err error) {

	resp, err := r.servicePager.NextPage(r.ctx)
	if err != nil {
		return response, err
	}

	for _, v := range resp.Segment.BlobItems {
		response.Objects = append(response.Objects, implementable.DataStorageObject{
			Key:  *v.Name,
			Name: filepath.Base(*v.Name),
		})
	}

	if resp.NextMarker != nil && *resp.NextMarker != "" {
		response.Cursor = resp.NextMarker
	}

	return response, nil
}

func (r *pager) More() bool {
	return r.servicePager.More()
}
