package implementable

type ListObjectsInput struct {
	Prefix            *string
	Cursor            *string
	MaxResultsPerPage *int
}

type DataStorageObject struct {
	Key  string
	Name string
}

type ListObjectsResponse struct {
	Objects []DataStorageObject
	Cursor  *string
}

type ListObjectsPager Pager[interface{}, ListObjectsResponse]

type DataStorageClient interface {
	UploadFile(srcAbsPath string, remoteObjectKey string) (interface{}, error)
	DownloadFile(remoteObjectKey string, destAbsPath string) (interface{}, error)
	NewListObjectsPager(input *ListObjectsInput) ListObjectsPager
}

type CloudClient interface {
	NewDataStorageClient(dataStorageId string, dataStorageContainerId string) (DataStorageClient, error)
}
