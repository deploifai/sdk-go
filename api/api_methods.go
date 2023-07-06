package api

import (
	"errors"
	"fmt"
	"github.com/Yamashou/gqlgenc/clientv2"
)

func (api *API[T]) GetClient() T {
	return api.Client
}

func (api *API[T]) ProcessError(err error) error {
	if handledError, ok := err.(*clientv2.ErrorResponse); ok {
		msg := "handled error: "
		if handledError.NetworkError != nil {
			msg = msg + fmt.Sprintf("network error: [status code = %d] %s\n", handledError.NetworkError.Code, handledError.NetworkError.Message)
		} else {
			msg = msg + fmt.Sprintf("graphql error: %v\n", handledError.GqlErrors)
		}
		return errors.New(msg)
	}

	return errors.New(fmt.Sprintf("unhandled error: %s\n", err.Error()))
}
