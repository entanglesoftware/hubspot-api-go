package decorator

import (
	"fmt"
	"time"
)

// RetryDecorator struct to handle retry logic
type RetryDecorator struct {
	TenSecondlyRolling     string
	SecondlyLimitMessage   string
	RetryTimeout           map[int]time.Duration
	NumberOfApiCallRetries int
}

// NewRetryDecorator initializes a RetryDecorator with specified retries
func NewRetryDecorator(numberOfApiCallRetries int) *RetryDecorator {
	return &RetryDecorator{
		TenSecondlyRolling:   "TEN_SECONDLY_ROLLING",
		SecondlyLimitMessage: "You have reached your secondly limit.",
		RetryTimeout: map[int]time.Duration{
			500: 200 * time.Millisecond, // Internal Server Error
			429: 10 * time.Second,       // Too Many Requests
			999: 1 * time.Second,        // Too Many Search Requests
		},
		NumberOfApiCallRetries: numberOfApiCallRetries,
	}
}

// Decorate wraps a method with retry logic
func (d *RetryDecorator) Decorate(method func(args ...interface{}) (interface{}, error)) func(args ...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		var resultSuccess interface{}
		var resultRejected error

		for index := 1; index <= d.NumberOfApiCallRetries; index++ {
			result, err := method(args...)
			if err == nil {
				resultSuccess = result
				resultRejected = nil
				break
			}

			resultRejected = err
			statusCode := d.getStatusCode(err)

			if index == d.NumberOfApiCallRetries {
				break
			}

			switch statusCode {
			case 500:
				d.waitAfterRequestFailure(statusCode, index, d.RetryTimeout[500])
			case 429:
				d.waitAfterRequestFailure(statusCode, index, d.RetryTimeout[429])
			default:
				break
			}
		}

		if resultRejected != nil {
			return nil, resultRejected
		}
		return resultSuccess, nil
	}
}

// getStatusCode simulates error code extraction from an error
func (d *RetryDecorator) getStatusCode(err error) int {
	// Placeholder - adapt to your error struct to get the actual code
	return 500 // Default to 500 for example purposes
}

// waitAfterRequestFailure waits for a retry timeout after a failed request
func (d *RetryDecorator) waitAfterRequestFailure(statusCode int, retryNumber int, retryTimeout time.Duration) {
	fmt.Printf("Request failed with status code [%d], will retry [%d] time(s) in [%v] ms\n", statusCode, retryNumber, retryTimeout)
	time.Sleep(retryTimeout * time.Duration(retryNumber))
}
