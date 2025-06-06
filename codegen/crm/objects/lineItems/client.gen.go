// Package lineItems provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package lineItems

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
	// GetLineItems request
	GetLineItems(ctx context.Context, params *GetLineItemsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateLineItemWithBody request with any body
	CreateLineItemWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateLineItem(ctx context.Context, body CreateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BatchCreateLineItemsWithBody request with any body
	BatchCreateLineItemsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BatchCreateLineItems(ctx context.Context, body BatchCreateLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SearchLineItemsWithBody request with any body
	SearchLineItemsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SearchLineItems(ctx context.Context, body SearchLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteLineItemById request
	DeleteLineItemById(ctx context.Context, lineItemId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLineItemById request
	GetLineItemById(ctx context.Context, lineItemId string, params *GetLineItemByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateLineItemWithBody request with any body
	UpdateLineItemWithBody(ctx context.Context, lineItemId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateLineItem(ctx context.Context, lineItemId string, body UpdateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetLineItems(ctx context.Context, params *GetLineItemsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLineItemsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateLineItemWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateLineItemRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateLineItem(ctx context.Context, body CreateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateLineItemRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BatchCreateLineItemsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBatchCreateLineItemsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BatchCreateLineItems(ctx context.Context, body BatchCreateLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBatchCreateLineItemsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchLineItemsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchLineItemsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchLineItems(ctx context.Context, body SearchLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchLineItemsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteLineItemById(ctx context.Context, lineItemId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteLineItemByIdRequest(c.Server, lineItemId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetLineItemById(ctx context.Context, lineItemId string, params *GetLineItemByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLineItemByIdRequest(c.Server, lineItemId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateLineItemWithBody(ctx context.Context, lineItemId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateLineItemRequestWithBody(c.Server, lineItemId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateLineItem(ctx context.Context, lineItemId string, body UpdateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateLineItemRequest(c.Server, lineItemId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetLineItemsRequest generates requests for GetLineItems
func NewGetLineItemsRequest(server string, params *GetLineItemsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/line_items")
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

// NewCreateLineItemRequest calls the generic CreateLineItem builder with application/json body
func NewCreateLineItemRequest(server string, body CreateLineItemJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateLineItemRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateLineItemRequestWithBody generates requests for CreateLineItem with any type of body
func NewCreateLineItemRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/line_items")
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

// NewBatchCreateLineItemsRequest calls the generic BatchCreateLineItems builder with application/json body
func NewBatchCreateLineItemsRequest(server string, body BatchCreateLineItemsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBatchCreateLineItemsRequestWithBody(server, "application/json", bodyReader)
}

// NewBatchCreateLineItemsRequestWithBody generates requests for BatchCreateLineItems with any type of body
func NewBatchCreateLineItemsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/line_items/batch/create")
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

// NewSearchLineItemsRequest calls the generic SearchLineItems builder with application/json body
func NewSearchLineItemsRequest(server string, body SearchLineItemsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSearchLineItemsRequestWithBody(server, "application/json", bodyReader)
}

// NewSearchLineItemsRequestWithBody generates requests for SearchLineItems with any type of body
func NewSearchLineItemsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/line_items/search")
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

// NewDeleteLineItemByIdRequest generates requests for DeleteLineItemById
func NewDeleteLineItemByIdRequest(server string, lineItemId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "lineItemId", runtime.ParamLocationPath, lineItemId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/line_items/%s", pathParam0)
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

// NewGetLineItemByIdRequest generates requests for GetLineItemById
func NewGetLineItemByIdRequest(server string, lineItemId string, params *GetLineItemByIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "lineItemId", runtime.ParamLocationPath, lineItemId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/line_items/%s", pathParam0)
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

// NewUpdateLineItemRequest calls the generic UpdateLineItem builder with application/json body
func NewUpdateLineItemRequest(server string, lineItemId string, body UpdateLineItemJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateLineItemRequestWithBody(server, lineItemId, "application/json", bodyReader)
}

// NewUpdateLineItemRequestWithBody generates requests for UpdateLineItem with any type of body
func NewUpdateLineItemRequestWithBody(server string, lineItemId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "lineItemId", runtime.ParamLocationPath, lineItemId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/line_items/%s", pathParam0)
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
	// GetLineItemsWithResponse request
	GetLineItemsWithResponse(ctx context.Context, params *GetLineItemsParams, reqEditors ...RequestEditorFn) (*GetLineItemsResponse, error)

	// CreateLineItemWithBodyWithResponse request with any body
	CreateLineItemWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateLineItemResponse, error)

	CreateLineItemWithResponse(ctx context.Context, body CreateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateLineItemResponse, error)

	// BatchCreateLineItemsWithBodyWithResponse request with any body
	BatchCreateLineItemsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BatchCreateLineItemsResponse, error)

	BatchCreateLineItemsWithResponse(ctx context.Context, body BatchCreateLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*BatchCreateLineItemsResponse, error)

	// SearchLineItemsWithBodyWithResponse request with any body
	SearchLineItemsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchLineItemsResponse, error)

	SearchLineItemsWithResponse(ctx context.Context, body SearchLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchLineItemsResponse, error)

	// DeleteLineItemByIdWithResponse request
	DeleteLineItemByIdWithResponse(ctx context.Context, lineItemId string, reqEditors ...RequestEditorFn) (*DeleteLineItemByIdResponse, error)

	// GetLineItemByIdWithResponse request
	GetLineItemByIdWithResponse(ctx context.Context, lineItemId string, params *GetLineItemByIdParams, reqEditors ...RequestEditorFn) (*GetLineItemByIdResponse, error)

	// UpdateLineItemWithBodyWithResponse request with any body
	UpdateLineItemWithBodyWithResponse(ctx context.Context, lineItemId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLineItemResponse, error)

	UpdateLineItemWithResponse(ctx context.Context, lineItemId string, body UpdateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLineItemResponse, error)
}

type GetLineItemsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LineItemsResponse
}

// Status returns HTTPResponse.Status
func (r GetLineItemsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLineItemsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateLineItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the lineItem was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the lineItem was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the created lineItem.
		Id string `json:"id,omitempty"`

		// Properties Properties of the created lineItem.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the lineItem's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the lineItem was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
	JSON400 *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r CreateLineItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateLineItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BatchCreateLineItemsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *BatchResponseLineItems
	JSON400      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r BatchCreateLineItemsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BatchCreateLineItemsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SearchLineItemsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LineItemsResponse
}

// Status returns HTTPResponse.Status
func (r SearchLineItemsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchLineItemsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteLineItemByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteLineItemByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteLineItemByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetLineItemByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LineItemResponse
}

// Status returns HTTPResponse.Status
func (r GetLineItemByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLineItemByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateLineItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the lineItem was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the lineItem was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the updated lineItem.
		Id string `json:"id,omitempty"`

		// Properties Properties of the updated lineItem.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the lineItem's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the lineItem was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateLineItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateLineItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetLineItemsWithResponse request returning *GetLineItemsResponse
func (c *ClientWithResponses) GetLineItemsWithResponse(ctx context.Context, params *GetLineItemsParams, reqEditors ...RequestEditorFn) (*GetLineItemsResponse, error) {
	rsp, err := c.GetLineItems(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLineItemsResponse(rsp)
}

// CreateLineItemWithBodyWithResponse request with arbitrary body returning *CreateLineItemResponse
func (c *ClientWithResponses) CreateLineItemWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateLineItemResponse, error) {
	rsp, err := c.CreateLineItemWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateLineItemResponse(rsp)
}

func (c *ClientWithResponses) CreateLineItemWithResponse(ctx context.Context, body CreateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateLineItemResponse, error) {
	rsp, err := c.CreateLineItem(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateLineItemResponse(rsp)
}

// BatchCreateLineItemsWithBodyWithResponse request with arbitrary body returning *BatchCreateLineItemsResponse
func (c *ClientWithResponses) BatchCreateLineItemsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BatchCreateLineItemsResponse, error) {
	rsp, err := c.BatchCreateLineItemsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBatchCreateLineItemsResponse(rsp)
}

func (c *ClientWithResponses) BatchCreateLineItemsWithResponse(ctx context.Context, body BatchCreateLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*BatchCreateLineItemsResponse, error) {
	rsp, err := c.BatchCreateLineItems(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBatchCreateLineItemsResponse(rsp)
}

// SearchLineItemsWithBodyWithResponse request with arbitrary body returning *SearchLineItemsResponse
func (c *ClientWithResponses) SearchLineItemsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchLineItemsResponse, error) {
	rsp, err := c.SearchLineItemsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchLineItemsResponse(rsp)
}

func (c *ClientWithResponses) SearchLineItemsWithResponse(ctx context.Context, body SearchLineItemsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchLineItemsResponse, error) {
	rsp, err := c.SearchLineItems(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchLineItemsResponse(rsp)
}

// DeleteLineItemByIdWithResponse request returning *DeleteLineItemByIdResponse
func (c *ClientWithResponses) DeleteLineItemByIdWithResponse(ctx context.Context, lineItemId string, reqEditors ...RequestEditorFn) (*DeleteLineItemByIdResponse, error) {
	rsp, err := c.DeleteLineItemById(ctx, lineItemId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteLineItemByIdResponse(rsp)
}

// GetLineItemByIdWithResponse request returning *GetLineItemByIdResponse
func (c *ClientWithResponses) GetLineItemByIdWithResponse(ctx context.Context, lineItemId string, params *GetLineItemByIdParams, reqEditors ...RequestEditorFn) (*GetLineItemByIdResponse, error) {
	rsp, err := c.GetLineItemById(ctx, lineItemId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLineItemByIdResponse(rsp)
}

// UpdateLineItemWithBodyWithResponse request with arbitrary body returning *UpdateLineItemResponse
func (c *ClientWithResponses) UpdateLineItemWithBodyWithResponse(ctx context.Context, lineItemId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLineItemResponse, error) {
	rsp, err := c.UpdateLineItemWithBody(ctx, lineItemId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateLineItemResponse(rsp)
}

func (c *ClientWithResponses) UpdateLineItemWithResponse(ctx context.Context, lineItemId string, body UpdateLineItemJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLineItemResponse, error) {
	rsp, err := c.UpdateLineItem(ctx, lineItemId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateLineItemResponse(rsp)
}

// ParseGetLineItemsResponse parses an HTTP response from a GetLineItemsWithResponse call
func ParseGetLineItemsResponse(rsp *http.Response) (*GetLineItemsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLineItemsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LineItemsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateLineItemResponse parses an HTTP response from a CreateLineItemWithResponse call
func ParseCreateLineItemResponse(rsp *http.Response) (*CreateLineItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateLineItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the lineItem was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the lineItem was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the created lineItem.
			Id string `json:"id,omitempty"`

			// Properties Properties of the created lineItem.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the lineItem's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the lineItem was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseBatchCreateLineItemsResponse parses an HTTP response from a BatchCreateLineItemsWithResponse call
func ParseBatchCreateLineItemsResponse(rsp *http.Response) (*BatchCreateLineItemsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BatchCreateLineItemsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest BatchResponseLineItems
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseSearchLineItemsResponse parses an HTTP response from a SearchLineItemsWithResponse call
func ParseSearchLineItemsResponse(rsp *http.Response) (*SearchLineItemsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SearchLineItemsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LineItemsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteLineItemByIdResponse parses an HTTP response from a DeleteLineItemByIdWithResponse call
func ParseDeleteLineItemByIdResponse(rsp *http.Response) (*DeleteLineItemByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteLineItemByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetLineItemByIdResponse parses an HTTP response from a GetLineItemByIdWithResponse call
func ParseGetLineItemByIdResponse(rsp *http.Response) (*GetLineItemByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLineItemByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LineItemResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateLineItemResponse parses an HTTP response from a UpdateLineItemWithResponse call
func ParseUpdateLineItemResponse(rsp *http.Response) (*UpdateLineItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateLineItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the lineItem was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the lineItem was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the updated lineItem.
			Id string `json:"id,omitempty"`

			// Properties Properties of the updated lineItem.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the lineItem's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the lineItem was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
