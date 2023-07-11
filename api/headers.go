package api

// WithAuthHeader returns a RequestHeader with the key "Authorization" and the value of the authToken.
func WithAuthHeader(authToken string) RequestHeader {
	return RequestHeader{Key: "Authorization", Value: authToken}
}

// ContentType is an enum for the content type of a request.
type ContentType string

const (
	// ContentTypeJson is the content type for a request with a json body in utf-8.
	ContentTypeJson ContentType = "application/json; charset=UTF-8"
)

// WithContentType returns a RequestHeader with the key "Content-Type" and the string value of the contentType enum.
func WithContentType(contentType ContentType) RequestHeader {
	return RequestHeader{Key: "Content-Type", Value: string(contentType)}
}
