package implementable

type DataStorageClient interface {
	UploadFile(srcAbsPath string, remoteObjectKey string) (interface{}, error)
	DownloadFile(remoteObjectKey string, destAbsPath string) (interface{}, error)
}

type CloudClient interface {
	NewDataStorageClient(dataStorageId string, dataStorageContainerId string) (DataStorageClient, error)
}
