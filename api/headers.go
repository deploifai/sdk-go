package api

// WithAuthHeader returns a RequestHeader with the key "Authorization" and the value of the authToken
func WithAuthHeader(authToken string) RequestHeader {
	return RequestHeader{Key: "Authorization", Value: authToken}
}

type ContentType string

const (
	ContentTypeJson ContentType = "application/json; charset=UTF-8"
)

func WithContentType(contentType ContentType) RequestHeader {
	return RequestHeader{Key: "Content-Type", Value: string(contentType)}
}
