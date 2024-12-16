// Package contacts provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package contacts

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
	// GetContacts request
	GetContacts(ctx context.Context, params *GetContactsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateContactWithBody request with any body
	CreateContactWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateContact(ctx context.Context, body CreateContactJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetContactById request
	GetContactById(ctx context.Context, contactId int64, params *GetContactByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetContacts(ctx context.Context, params *GetContactsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetContactsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateContactWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateContactRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateContact(ctx context.Context, body CreateContactJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateContactRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetContactById(ctx context.Context, contactId int64, params *GetContactByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetContactByIdRequest(c.Server, contactId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetContactsRequest generates requests for GetContacts
func NewGetContactsRequest(server string, params *GetContactsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/contacts")
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

// NewCreateContactRequest calls the generic CreateContact builder with application/json body
func NewCreateContactRequest(server string, body CreateContactJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateContactRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateContactRequestWithBody generates requests for CreateContact with any type of body
func NewCreateContactRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/contacts")
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

// NewGetContactByIdRequest generates requests for GetContactById
func NewGetContactByIdRequest(server string, contactId int64, params *GetContactByIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "contactId", runtime.ParamLocationPath, contactId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/contacts/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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
	// GetContactsWithResponse request
	GetContactsWithResponse(ctx context.Context, params *GetContactsParams, reqEditors ...RequestEditorFn) (*GetContactsResponse, error)

	// CreateContactWithBodyWithResponse request with any body
	CreateContactWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateContactResponse, error)

	CreateContactWithResponse(ctx context.Context, body CreateContactJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateContactResponse, error)

	// GetContactByIdWithResponse request
	GetContactByIdWithResponse(ctx context.Context, contactId int64, params *GetContactByIdParams, reqEditors ...RequestEditorFn) (*GetContactByIdResponse, error)
}

type GetContactsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ContactsResponse
}

// Status returns HTTPResponse.Status
func (r GetContactsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetContactsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateContactResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// CreatedAt Timestamp when the contact was created.
		CreatedAt *time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the created contact.
		Id *string `json:"id,omitempty"`

		// Properties Properties of the created contact.
		Properties *map[string]interface{} `json:"properties,omitempty"`

		// UpdatedAt Timestamp when the contact was last updated.
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateContactResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateContactResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetContactByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ContactsResponse
}

// Status returns HTTPResponse.Status
func (r GetContactByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetContactByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetContactsWithResponse request returning *GetContactsResponse
func (c *ClientWithResponses) GetContactsWithResponse(ctx context.Context, params *GetContactsParams, reqEditors ...RequestEditorFn) (*GetContactsResponse, error) {
	rsp, err := c.GetContacts(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetContactsResponse(rsp)
}

// CreateContactWithBodyWithResponse request with arbitrary body returning *CreateContactResponse
func (c *ClientWithResponses) CreateContactWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateContactResponse, error) {
	rsp, err := c.CreateContactWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateContactResponse(rsp)
}

func (c *ClientWithResponses) CreateContactWithResponse(ctx context.Context, body CreateContactJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateContactResponse, error) {
	rsp, err := c.CreateContact(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateContactResponse(rsp)
}

// GetContactByIdWithResponse request returning *GetContactByIdResponse
func (c *ClientWithResponses) GetContactByIdWithResponse(ctx context.Context, contactId int64, params *GetContactByIdParams, reqEditors ...RequestEditorFn) (*GetContactByIdResponse, error) {
	rsp, err := c.GetContactById(ctx, contactId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetContactByIdResponse(rsp)
}

// ParseGetContactsResponse parses an HTTP response from a GetContactsWithResponse call
func ParseGetContactsResponse(rsp *http.Response) (*GetContactsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetContactsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContactsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateContactResponse parses an HTTP response from a CreateContactWithResponse call
func ParseCreateContactResponse(rsp *http.Response) (*CreateContactResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateContactResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// CreatedAt Timestamp when the contact was created.
			CreatedAt *time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the created contact.
			Id *string `json:"id,omitempty"`

			// Properties Properties of the created contact.
			Properties *map[string]interface{} `json:"properties,omitempty"`

			// UpdatedAt Timestamp when the contact was last updated.
			UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseGetContactByIdResponse parses an HTTP response from a GetContactByIdWithResponse call
func ParseGetContactByIdResponse(rsp *http.Response) (*GetContactByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetContactByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContactsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
