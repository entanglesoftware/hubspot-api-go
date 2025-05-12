package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"net/http"
	"net/url"
)

// Headers type for headers map
type Headers map[string]string

// Request struct for HTTP request
type Request struct {
	opts    Options
	config  configuration.Configuration
	baseURL string
	url     *url.URL
	method  string
	Headers Headers
	body    []byte
}

// NewHttpRequest NewRequest initializes a new request
func NewHttpRequest(config configuration.Configuration, opts Options) (*Request, error) {
	req := &Request{
		config:  config,
		opts:    opts,
		Headers: Headers{},
	}

	if config.BasePath != "" {
		req.baseURL = config.BasePath
	} else {
		req.baseURL = "https://api.hubapi.com"
	}

	err := req.init()
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Initialize request properties
func (r *Request) init() error {
	r.url = r.generateUrl()
	r.method = r.opts.Method
	if r.method == "" {
		r.method = "GET"
	}
	r.initHeaders()
	err := r.applyAuth()
	if err != nil {
		return err
	}
	err = r.setBody()
	if err != nil {
		return err
	}
	return nil
}

// Generate the URL based on options
func (r *Request) generateUrl() *url.URL {
	urlStr := r.opts.OverlapUrl
	if urlStr == "" {
		urlStr = r.baseURL + r.opts.Path
	}

	parsedUrl, _ := url.Parse(urlStr)

	// Add query string parameters if provided
	if r.opts.QS != nil {
		q := parsedUrl.Query()
		for key, value := range r.opts.QS {
			q.Set(key, value)
		}
		parsedUrl.RawQuery = q.Encode()
	}

	return parsedUrl
}

// Apply authentication based on the configuration
func (r *Request) applyAuth() error {
	if r.config.APIKey != "" {
		query := r.url.Query()
		query.Set("hapikey", r.config.APIKey)
		r.url.RawQuery = query.Encode()
	} else if r.config.AccessToken != "" {
		r.Headers["Authorization"] = "Bearer " + r.config.AccessToken
	} else {
		return errors.New("authentication method not set in configuration")
	}
	return nil
}

// Initialize headers based on options and configuration
func (r *Request) initHeaders() {
	r.Headers = Headers{
		"Content-Type": "application/json",
	}

	for k, v := range r.config.DefaultHeaders {
		r.Headers[k] = v
	}

	for k, v := range r.opts.Headers {
		r.Headers[k] = v
	}
}

// Set the request body if provided in options
func (r *Request) setBody() error {
	if r.opts.Body != nil {
		var body []byte
		var err error

		if r.opts.DefaultJSON {
			body, err = json.Marshal(r.opts.Body)
			if err != nil {
				return fmt.Errorf("error marshalling body to JSON: %w", err)
			}
		} else {
			body = []byte(fmt.Sprint(r.opts.Body))
		}
		r.body = body
	}
	return nil
}

// GetSendData prepares the data to be sent in the request
func (r *Request) GetSendData() (*http.Request, error) {
	req, err := http.NewRequest(r.method, r.url.String(), bytes.NewBuffer(r.body))
	if err != nil {
		return nil, err
	}

	// Set headers for the request
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}
