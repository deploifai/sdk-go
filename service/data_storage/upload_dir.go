package data_storage

import (
	"context"
	"fmt"
	"github.com/deploifai/sdk-go/api/generated"
	"github.com/panjf2000/ants/v2"
	"io/fs"
	"path/filepath"
	"runtime"
	"sync"
)

type UploadDirInput struct {
	srcAbsPath      string
	remoteObjectKey string
}

type UploadDirOptions struct {
	// Concurrency is the number of workers to execute.
	// Default is number of CPUs available.
	Concurrency *int
}

type task struct {
	filePath   string
	resultChan chan<- interface{}
	errChan    chan<- error
}

// UploadDir uploads a local directory to a remote directory in the data storage container.
// This function is meant to be used as a goroutine.
func (c *Client) UploadDir(
	ctx context.Context,
	where generated.DataStorageContainerWhereUniqueInput,
	data UploadDirInput,
	fileCountChan chan<- int,
	resultChan chan<- interface{},
	options *UploadDirOptions) error {

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

	// clean the srcAbsPath and remoteObjectKey
	srcAbsPath := filepath.Clean(data.srcAbsPath) + "/"
	remoteObjectKey := filepath.ToSlash(filepath.Clean(data.remoteObjectKey))
	// remove leading slash if any
	if len(remoteObjectKey) > 0 && remoteObjectKey[0] == '/' {
		remoteObjectKey = remoteObjectKey[1:]
	}
	// if remoteObjectKey is at root, then it should just be an empty string
	// if remoteObjectKey is not at root, then it should end with a slash
	if remoteObjectKey == "." || remoteObjectKey == "" {
		remoteObjectKey = ""
	} else {
		remoteObjectKey = remoteObjectKey + "/"
	}
	fmt.Println("root remote object key:", remoteObjectKey)

	filePaths, err := listFiles(srcAbsPath)
	if err != nil {
		return err
	}

	// inform number of files via fileCountChan
	fileCountChan <- len(filePaths)

	var wg sync.WaitGroup
	var errChan = make(chan error, len(filePaths))

	pool, err := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		task := i.(task)
		remoteObjectKey := remoteObjectKey + filepath.ToSlash(task.filePath[len(srcAbsPath):])
		fmt.Println("remote object key:", remoteObjectKey)
		result, err := dataStorageClient.UploadFile(task.filePath, remoteObjectKey)
		if err != nil {
			task.errChan <- err
		} else {
			task.resultChan <- result
		}
		wg.Done()
	})
	if err != nil {
		return err
	}
	defer pool.Release()

	for _, filePath := range filePaths {
		wg.Add(1)
		err := pool.Invoke(task{filePath: filePath, resultChan: resultChan, errChan: errChan})
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

func listFiles(root string) (files []string, err error) {

	err = filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
