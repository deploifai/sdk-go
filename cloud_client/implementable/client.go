package implementable

type ListObjectsInput struct {
	Prefix     *string
	MaxResults *int
	Cursor     *string
}

type DataStorageObject struct {
	Key  string
	Name string
}

type ListObjectsResponse struct {
	Objects []DataStorageObject
	Cursor  *string
}

type DataStorageClient interface {
	UploadFile(srcAbsPath string, remoteObjectKey string) (interface{}, error)
	DownloadFile(remoteObjectKey string, destAbsPath string) (interface{}, error)
	ListObjects(input *ListObjectsInput) (ListObjectsResponse, error)
}

type CloudClient interface {
	NewDataStorageClient(dataStorageId string, dataStorageContainerId string) (DataStorageClient, error)
}
