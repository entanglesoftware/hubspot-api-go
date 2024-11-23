package http

import (
	"fmt"
	"io"
	"net/http"
)

// Client struct for HTTP client
type Client struct{}

// NewHttpClient initializes and returns a new Client instance
func NewHttpClient() *Client {
	return &Client{}
}

// SendAdapter wraps Send to match the generalized IDecorator signature
func (h *Client) SendAdapter(args ...interface{}) (interface{}, error) {
	// Ensure the first argument is of type *Request
	if len(args) < 1 {
		return nil, fmt.Errorf("missing arguments in SendAdapter")
	}
	request, ok := args[0].(*Request)
	if !ok {
		return nil, fmt.Errorf("expected *Request, got %T", args[0])
	}

	// Call the original Send method
	return h.Send(request)
}

// Send executes the request and returns the response body or an error
func (h *Client) Send(request *Request) (*Response, error) {
	httpRequest, err := request.GetSendData()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:       body,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
	}, nil
}
