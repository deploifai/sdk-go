package utils

import (
	"fmt"
	"testing"
)

// TestSuccessCallWithRetries tests the function CallWithRetries
// with a function that returns an error 2 times and then returns a result
func TestSuccessCallWithRetries(t *testing.T) {
	counter := 0
	f := func() (int32, error) {
		if counter < 2 {
			counter++
			return 0, fmt.Errorf("error")
		} else {
			return 1, nil
		}
	}

	result, err := CallWithRetries(f, nil, nil)
	if err != nil {
		t.Errorf("CallWithRetries returned an error after retrying")
	}
	if result != 1 {
		t.Errorf("CallWithRetries returned a wrong result")
	}
}

// TestErrorCallWithRetries tests the function CallWithRetries
// with a function that returns an error after 2 retries
func TestErrorCallWithRetries(t *testing.T) {
	counter := 0
	f := func() (int32, error) {
		if counter < 2 {
			counter++
			return 0, fmt.Errorf("error")
		} else {
			return 1, nil
		}
	}

	retryCount := 2
	_, err := CallWithRetries(f, &retryCount, nil)
	if err == nil {
		t.Errorf("CallWithRetries did not return an error after retrying 3 times")
	}
}
