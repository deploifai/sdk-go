package utils

import "time"

// CallWithRetries calls a function with a number of retries and a retry interval
// retryCount default is 3, retryInterval default is 3 seconds
// returns the result and error of the function
func CallWithRetries[T any](fn func() (T, error), retryCount *int, retryInterval *time.Duration) (result T, err error) {

	// setup default values for retryCount
	if retryCount == nil {
		retryCount = new(int)
		*retryCount = 3
	}

	// setup default values for retryInterval
	if retryInterval == nil {
		retryInterval = new(time.Duration)
		*retryInterval = 3 * time.Second
	}

	// retry
	for i := 0; i < *retryCount; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		time.Sleep(*retryInterval)
	}
	return result, err
}
