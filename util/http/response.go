package http

import "net/http"

// Response HttpResponse struct to encapsulate the response details
type Response struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
}
