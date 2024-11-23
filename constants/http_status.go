package constants

type StatusCode int

const (
	TooManyRequests StatusCode = 429
	MinServerError  StatusCode = 500
	MaxServerError  StatusCode = 599
)
