// Package quotes provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package quotes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetQuotes request
	GetQuotes(ctx context.Context, params *GetQuotesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateQuoteWithBody request with any body
	CreateQuoteWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateQuote(ctx context.Context, body CreateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SearchQuotesWithBody request with any body
	SearchQuotesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SearchQuotes(ctx context.Context, body SearchQuotesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteQuoteById request
	DeleteQuoteById(ctx context.Context, quoteId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetQuoteById request
	GetQuoteById(ctx context.Context, quoteId int64, params *GetQuoteByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateQuoteWithBody request with any body
	UpdateQuoteWithBody(ctx context.Context, quoteId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateQuote(ctx context.Context, quoteId string, body UpdateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetQuotes(ctx context.Context, params *GetQuotesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetQuotesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateQuoteWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateQuoteRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateQuote(ctx context.Context, body CreateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateQuoteRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchQuotesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchQuotesRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchQuotes(ctx context.Context, body SearchQuotesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchQuotesRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteQuoteById(ctx context.Context, quoteId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteQuoteByIdRequest(c.Server, quoteId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetQuoteById(ctx context.Context, quoteId int64, params *GetQuoteByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetQuoteByIdRequest(c.Server, quoteId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateQuoteWithBody(ctx context.Context, quoteId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateQuoteRequestWithBody(c.Server, quoteId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateQuote(ctx context.Context, quoteId string, body UpdateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateQuoteRequest(c.Server, quoteId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetQuotesRequest generates requests for GetQuotes
func NewGetQuotesRequest(server string, params *GetQuotesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Limit != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.After != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "after", runtime.ParamLocationQuery, *params.After); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Properties != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "properties", runtime.ParamLocationQuery, *params.Properties); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PropertiesWithHistory != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "propertiesWithHistory", runtime.ParamLocationQuery, *params.PropertiesWithHistory); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Associations != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "associations", runtime.ParamLocationQuery, *params.Associations); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Archived != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "archived", runtime.ParamLocationQuery, *params.Archived); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateQuoteRequest calls the generic CreateQuote builder with application/json body
func NewCreateQuoteRequest(server string, body CreateQuoteJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateQuoteRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateQuoteRequestWithBody generates requests for CreateQuote with any type of body
func NewCreateQuoteRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewSearchQuotesRequest calls the generic SearchQuotes builder with application/json body
func NewSearchQuotesRequest(server string, body SearchQuotesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSearchQuotesRequestWithBody(server, "application/json", bodyReader)
}

// NewSearchQuotesRequestWithBody generates requests for SearchQuotes with any type of body
func NewSearchQuotesRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/search")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteQuoteByIdRequest generates requests for DeleteQuoteById
func NewDeleteQuoteByIdRequest(server string, quoteId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "quoteId", runtime.ParamLocationPath, quoteId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetQuoteByIdRequest generates requests for GetQuoteById
func NewGetQuoteByIdRequest(server string, quoteId int64, params *GetQuoteByIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "quoteId", runtime.ParamLocationPath, quoteId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.IdProperty != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "idProperty", runtime.ParamLocationQuery, *params.IdProperty); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Properties != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "properties", runtime.ParamLocationQuery, *params.Properties); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PropertiesWithHistory != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "propertiesWithHistory", runtime.ParamLocationQuery, *params.PropertiesWithHistory); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Associations != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "associations", runtime.ParamLocationQuery, *params.Associations); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Archived != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "archived", runtime.ParamLocationQuery, *params.Archived); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateQuoteRequest calls the generic UpdateQuote builder with application/json body
func NewUpdateQuoteRequest(server string, quoteId string, body UpdateQuoteJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateQuoteRequestWithBody(server, quoteId, "application/json", bodyReader)
}

// NewUpdateQuoteRequestWithBody generates requests for UpdateQuote with any type of body
func NewUpdateQuoteRequestWithBody(server string, quoteId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "quoteId", runtime.ParamLocationPath, quoteId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetQuotesWithResponse request
	GetQuotesWithResponse(ctx context.Context, params *GetQuotesParams, reqEditors ...RequestEditorFn) (*GetQuotesResponse, error)

	// CreateQuoteWithBodyWithResponse request with any body
	CreateQuoteWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateQuoteResponse, error)

	CreateQuoteWithResponse(ctx context.Context, body CreateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateQuoteResponse, error)

	// SearchQuotesWithBodyWithResponse request with any body
	SearchQuotesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchQuotesResponse, error)

	SearchQuotesWithResponse(ctx context.Context, body SearchQuotesJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchQuotesResponse, error)

	// DeleteQuoteByIdWithResponse request
	DeleteQuoteByIdWithResponse(ctx context.Context, quoteId string, reqEditors ...RequestEditorFn) (*DeleteQuoteByIdResponse, error)

	// GetQuoteByIdWithResponse request
	GetQuoteByIdWithResponse(ctx context.Context, quoteId int64, params *GetQuoteByIdParams, reqEditors ...RequestEditorFn) (*GetQuoteByIdResponse, error)

	// UpdateQuoteWithBodyWithResponse request with any body
	UpdateQuoteWithBodyWithResponse(ctx context.Context, quoteId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateQuoteResponse, error)

	UpdateQuoteWithResponse(ctx context.Context, quoteId string, body UpdateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateQuoteResponse, error)
}

type GetQuotesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *QuotesResponse
}

// Status returns HTTPResponse.Status
func (r GetQuotesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetQuotesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateQuoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the quote was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the quote was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the created quote.
		Id string `json:"id,omitempty"`

		// Properties Properties of the created quote.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the quote's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the quote was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateQuoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateQuoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SearchQuotesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *QuotesResponse
}

// Status returns HTTPResponse.Status
func (r SearchQuotesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchQuotesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteQuoteByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteQuoteByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteQuoteByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetQuoteByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *QuoteResponse
}

// Status returns HTTPResponse.Status
func (r GetQuoteByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetQuoteByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateQuoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the quote was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the quote was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the updated quote.
		Id string `json:"id,omitempty"`

		// Properties Properties of the updated quote.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the quote's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the quote was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateQuoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateQuoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetQuotesWithResponse request returning *GetQuotesResponse
func (c *ClientWithResponses) GetQuotesWithResponse(ctx context.Context, params *GetQuotesParams, reqEditors ...RequestEditorFn) (*GetQuotesResponse, error) {
	rsp, err := c.GetQuotes(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetQuotesResponse(rsp)
}

// CreateQuoteWithBodyWithResponse request with arbitrary body returning *CreateQuoteResponse
func (c *ClientWithResponses) CreateQuoteWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateQuoteResponse, error) {
	rsp, err := c.CreateQuoteWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateQuoteResponse(rsp)
}

func (c *ClientWithResponses) CreateQuoteWithResponse(ctx context.Context, body CreateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateQuoteResponse, error) {
	rsp, err := c.CreateQuote(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateQuoteResponse(rsp)
}

// SearchQuotesWithBodyWithResponse request with arbitrary body returning *SearchQuotesResponse
func (c *ClientWithResponses) SearchQuotesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchQuotesResponse, error) {
	rsp, err := c.SearchQuotesWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchQuotesResponse(rsp)
}

func (c *ClientWithResponses) SearchQuotesWithResponse(ctx context.Context, body SearchQuotesJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchQuotesResponse, error) {
	rsp, err := c.SearchQuotes(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchQuotesResponse(rsp)
}

// DeleteQuoteByIdWithResponse request returning *DeleteQuoteByIdResponse
func (c *ClientWithResponses) DeleteQuoteByIdWithResponse(ctx context.Context, quoteId string, reqEditors ...RequestEditorFn) (*DeleteQuoteByIdResponse, error) {
	rsp, err := c.DeleteQuoteById(ctx, quoteId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteQuoteByIdResponse(rsp)
}

// GetQuoteByIdWithResponse request returning *GetQuoteByIdResponse
func (c *ClientWithResponses) GetQuoteByIdWithResponse(ctx context.Context, quoteId int64, params *GetQuoteByIdParams, reqEditors ...RequestEditorFn) (*GetQuoteByIdResponse, error) {
	rsp, err := c.GetQuoteById(ctx, quoteId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetQuoteByIdResponse(rsp)
}

// UpdateQuoteWithBodyWithResponse request with arbitrary body returning *UpdateQuoteResponse
func (c *ClientWithResponses) UpdateQuoteWithBodyWithResponse(ctx context.Context, quoteId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateQuoteResponse, error) {
	rsp, err := c.UpdateQuoteWithBody(ctx, quoteId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateQuoteResponse(rsp)
}

func (c *ClientWithResponses) UpdateQuoteWithResponse(ctx context.Context, quoteId string, body UpdateQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateQuoteResponse, error) {
	rsp, err := c.UpdateQuote(ctx, quoteId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateQuoteResponse(rsp)
}

// ParseGetQuotesResponse parses an HTTP response from a GetQuotesWithResponse call
func ParseGetQuotesResponse(rsp *http.Response) (*GetQuotesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetQuotesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest QuotesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateQuoteResponse parses an HTTP response from a CreateQuoteWithResponse call
func ParseCreateQuoteResponse(rsp *http.Response) (*CreateQuoteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateQuoteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the quote was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the quote was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the created quote.
			Id string `json:"id,omitempty"`

			// Properties Properties of the created quote.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the quote's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the quote was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseSearchQuotesResponse parses an HTTP response from a SearchQuotesWithResponse call
func ParseSearchQuotesResponse(rsp *http.Response) (*SearchQuotesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SearchQuotesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest QuotesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteQuoteByIdResponse parses an HTTP response from a DeleteQuoteByIdWithResponse call
func ParseDeleteQuoteByIdResponse(rsp *http.Response) (*DeleteQuoteByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteQuoteByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetQuoteByIdResponse parses an HTTP response from a GetQuoteByIdWithResponse call
func ParseGetQuoteByIdResponse(rsp *http.Response) (*GetQuoteByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetQuoteByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest QuoteResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateQuoteResponse parses an HTTP response from a UpdateQuoteWithResponse call
func ParseUpdateQuoteResponse(rsp *http.Response) (*UpdateQuoteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateQuoteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the quote was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the quote was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the updated quote.
			Id string `json:"id,omitempty"`

			// Properties Properties of the updated quote.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the quote's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the quote was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
