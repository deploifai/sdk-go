package dataset

import (
	"context"
	"fmt"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/deploifai/sdk-go/cloud_client/implementable"
	"github.com/panjf2000/ants/v2"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type DownloadDirInput struct {
	remoteObjectPrefix string
	destAbsPath        string
}

type DownloadDirOptions struct {
	// Concurrency is the number of workers to execute.
	// Default is number of CPUs available.
	Concurrency *int
}

func (c *Client) DownloadDir(
	ctx context.Context,
	where generated.DataStorageWhereUniqueInput,
	data DownloadDirInput,
	fileCountChan chan<- int,
	resultChan chan<- interface{},
	options *DownloadDirOptions) error {

	// declare options
	poolSize := runtime.NumCPU()

	// amend options
	if options != nil {
		if options.Concurrency != nil {
			poolSize = *options.Concurrency
		}
	}

	dataStorage, dataStorageContainer, err := getDataStorageAndContainer(ctx, c.options.API, where)
	if err != nil {
		return err
	}

	dataStorageClient, err := newDataStorageClient(ctx, c.options.API, dataStorage, dataStorageContainer)
	if err != nil {
		return err
	}

	prefix := cleanRemoteObjectPrefix(data.remoteObjectPrefix)
	fmt.Println("prefix: ", prefix)

	objects, err := listObjects(dataStorageClient, &implementable.ListObjectsInput{
		Prefix: &prefix,
	})
	if err != nil {
		return err
	}

	fmt.Println("number of objects: ", len(objects))

	fileCountChan <- len(objects)

	destAbsPath := filepath.Clean(data.destAbsPath) + "/"

	var wg sync.WaitGroup
	var errChan = make(chan error)

	pool, err := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		defer wg.Done()

		task := i.(workerTask)
		filePath := getDestAbsPath(destAbsPath, prefix, task.path)

		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			task.errChan <- err
			return
		}
		if _, err := os.Create(filePath); err != nil {
			task.errChan <- err
			return
		}

		result, err := dataStorageClient.DownloadFile(task.path, filePath)
		if err != nil {
			task.errChan <- err
		} else {
			task.resultChan <- result
		}
	})
	if err != nil {
		return err
	}
	defer pool.Release()

	for _, object := range objects {
		wg.Add(1)
		err := pool.Invoke(workerTask{path: object.Key, resultChan: resultChan, errChan: errChan})
		if err != nil {
			return err
		}
	}

	wg.Wait()

	// return the first error if any
	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

func listObjects(dataStorageClient implementable.DataStorageClient, input *implementable.ListObjectsInput) (objects []implementable.DataStorageObject, err error) {

	pager := dataStorageClient.NewListObjectsPager(input)

	for pager.More() {
		response, err := pager.NextPage(nil)
		if err != nil {
			return objects, err
		}

		objects = append(objects, response.Objects...)
	}

	return objects, nil
}

func getDestAbsPath(destAbsPath string, remoteObjectPrefix string, remoteObjectKey string) string {
	return destAbsPath + remoteObjectKey[len(remoteObjectPrefix):]
}
