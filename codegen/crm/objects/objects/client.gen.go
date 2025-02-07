// Package objects provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package objects

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
	// GetObjects request
	GetObjects(ctx context.Context, objectType string, params *GetObjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateObjectWithBody request with any body
	CreateObjectWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateObject(ctx context.Context, objectType string, body CreateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SearchObjectsWithBody request with any body
	SearchObjectsWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SearchObjects(ctx context.Context, objectType string, body SearchObjectsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteObject request
	DeleteObject(ctx context.Context, objectType string, objectId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetObjectByTypeAndId request
	GetObjectByTypeAndId(ctx context.Context, objectType string, objectId string, params *GetObjectByTypeAndIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateObjectWithBody request with any body
	UpdateObjectWithBody(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateObject(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, body UpdateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetObjects(ctx context.Context, objectType string, params *GetObjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetObjectsRequest(c.Server, objectType, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateObjectWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateObjectRequestWithBody(c.Server, objectType, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateObject(ctx context.Context, objectType string, body CreateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateObjectRequest(c.Server, objectType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchObjectsWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchObjectsRequestWithBody(c.Server, objectType, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchObjects(ctx context.Context, objectType string, body SearchObjectsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchObjectsRequest(c.Server, objectType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteObject(ctx context.Context, objectType string, objectId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteObjectRequest(c.Server, objectType, objectId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetObjectByTypeAndId(ctx context.Context, objectType string, objectId string, params *GetObjectByTypeAndIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetObjectByTypeAndIdRequest(c.Server, objectType, objectId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateObjectWithBody(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateObjectRequestWithBody(c.Server, objectType, objectId, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateObject(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, body UpdateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateObjectRequest(c.Server, objectType, objectId, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetObjectsRequest generates requests for GetObjects
func NewGetObjectsRequest(server string, objectType string, params *GetObjectsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/%s", pathParam0)
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

// NewCreateObjectRequest calls the generic CreateObject builder with application/json body
func NewCreateObjectRequest(server string, objectType string, body CreateObjectJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateObjectRequestWithBody(server, objectType, "application/json", bodyReader)
}

// NewCreateObjectRequestWithBody generates requests for CreateObject with any type of body
func NewCreateObjectRequestWithBody(server string, objectType string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/%s", pathParam0)
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

// NewSearchObjectsRequest calls the generic SearchObjects builder with application/json body
func NewSearchObjectsRequest(server string, objectType string, body SearchObjectsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSearchObjectsRequestWithBody(server, objectType, "application/json", bodyReader)
}

// NewSearchObjectsRequestWithBody generates requests for SearchObjects with any type of body
func NewSearchObjectsRequestWithBody(server string, objectType string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/%s/search", pathParam0)
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

// NewDeleteObjectRequest generates requests for DeleteObject
func NewDeleteObjectRequest(server string, objectType string, objectId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "objectId", runtime.ParamLocationPath, objectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/%s/%s", pathParam0, pathParam1)
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

// NewGetObjectByTypeAndIdRequest generates requests for GetObjectByTypeAndId
func NewGetObjectByTypeAndIdRequest(server string, objectType string, objectId string, params *GetObjectByTypeAndIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "objectId", runtime.ParamLocationPath, objectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/%s/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.PropertiesWithHistory != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "propertiesWithHistory", runtime.ParamLocationQuery, *params.PropertiesWithHistory); err != nil {
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

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "associations", runtime.ParamLocationQuery, *params.Associations); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateObjectRequest calls the generic UpdateObject builder with application/json body
func NewUpdateObjectRequest(server string, objectType string, objectId string, params *UpdateObjectParams, body UpdateObjectJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateObjectRequestWithBody(server, objectType, objectId, params, "application/json", bodyReader)
}

// NewUpdateObjectRequestWithBody generates requests for UpdateObject with any type of body
func NewUpdateObjectRequestWithBody(server string, objectType string, objectId string, params *UpdateObjectParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "objectId", runtime.ParamLocationPath, objectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/%s/%s", pathParam0, pathParam1)
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

		queryURL.RawQuery = queryValues.Encode()
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
	// GetObjectsWithResponse request
	GetObjectsWithResponse(ctx context.Context, objectType string, params *GetObjectsParams, reqEditors ...RequestEditorFn) (*GetObjectsResponse, error)

	// CreateObjectWithBodyWithResponse request with any body
	CreateObjectWithBodyWithResponse(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateObjectResponse, error)

	CreateObjectWithResponse(ctx context.Context, objectType string, body CreateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateObjectResponse, error)

	// SearchObjectsWithBodyWithResponse request with any body
	SearchObjectsWithBodyWithResponse(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchObjectsResponse, error)

	SearchObjectsWithResponse(ctx context.Context, objectType string, body SearchObjectsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchObjectsResponse, error)

	// DeleteObjectWithResponse request
	DeleteObjectWithResponse(ctx context.Context, objectType string, objectId string, reqEditors ...RequestEditorFn) (*DeleteObjectResponse, error)

	// GetObjectByTypeAndIdWithResponse request
	GetObjectByTypeAndIdWithResponse(ctx context.Context, objectType string, objectId string, params *GetObjectByTypeAndIdParams, reqEditors ...RequestEditorFn) (*GetObjectByTypeAndIdResponse, error)

	// UpdateObjectWithBodyWithResponse request with any body
	UpdateObjectWithBodyWithResponse(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateObjectResponse, error)

	UpdateObjectWithResponse(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, body UpdateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateObjectResponse, error)
}

type GetObjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
}

// Status returns HTTPResponse.Status
func (r GetObjectsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetObjectsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateObjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ObjectResponse
}

// Status returns HTTPResponse.Status
func (r CreateObjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateObjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SearchObjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ObjectsResponse
}

// Status returns HTTPResponse.Status
func (r SearchObjectsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchObjectsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteObjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteObjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteObjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetObjectByTypeAndIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ObjectResponse
}

// Status returns HTTPResponse.Status
func (r GetObjectByTypeAndIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetObjectByTypeAndIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateObjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the object was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the object was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the updated object.
		Id string `json:"id,omitempty"`

		// Properties Properties of the updated object.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the object's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the object was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateObjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateObjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetObjectsWithResponse request returning *GetObjectsResponse
func (c *ClientWithResponses) GetObjectsWithResponse(ctx context.Context, objectType string, params *GetObjectsParams, reqEditors ...RequestEditorFn) (*GetObjectsResponse, error) {
	rsp, err := c.GetObjects(ctx, objectType, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetObjectsResponse(rsp)
}

// CreateObjectWithBodyWithResponse request with arbitrary body returning *CreateObjectResponse
func (c *ClientWithResponses) CreateObjectWithBodyWithResponse(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateObjectResponse, error) {
	rsp, err := c.CreateObjectWithBody(ctx, objectType, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateObjectResponse(rsp)
}

func (c *ClientWithResponses) CreateObjectWithResponse(ctx context.Context, objectType string, body CreateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateObjectResponse, error) {
	rsp, err := c.CreateObject(ctx, objectType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateObjectResponse(rsp)
}

// SearchObjectsWithBodyWithResponse request with arbitrary body returning *SearchObjectsResponse
func (c *ClientWithResponses) SearchObjectsWithBodyWithResponse(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchObjectsResponse, error) {
	rsp, err := c.SearchObjectsWithBody(ctx, objectType, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchObjectsResponse(rsp)
}

func (c *ClientWithResponses) SearchObjectsWithResponse(ctx context.Context, objectType string, body SearchObjectsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchObjectsResponse, error) {
	rsp, err := c.SearchObjects(ctx, objectType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchObjectsResponse(rsp)
}

// DeleteObjectWithResponse request returning *DeleteObjectResponse
func (c *ClientWithResponses) DeleteObjectWithResponse(ctx context.Context, objectType string, objectId string, reqEditors ...RequestEditorFn) (*DeleteObjectResponse, error) {
	rsp, err := c.DeleteObject(ctx, objectType, objectId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteObjectResponse(rsp)
}

// GetObjectByTypeAndIdWithResponse request returning *GetObjectByTypeAndIdResponse
func (c *ClientWithResponses) GetObjectByTypeAndIdWithResponse(ctx context.Context, objectType string, objectId string, params *GetObjectByTypeAndIdParams, reqEditors ...RequestEditorFn) (*GetObjectByTypeAndIdResponse, error) {
	rsp, err := c.GetObjectByTypeAndId(ctx, objectType, objectId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetObjectByTypeAndIdResponse(rsp)
}

// UpdateObjectWithBodyWithResponse request with arbitrary body returning *UpdateObjectResponse
func (c *ClientWithResponses) UpdateObjectWithBodyWithResponse(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateObjectResponse, error) {
	rsp, err := c.UpdateObjectWithBody(ctx, objectType, objectId, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateObjectResponse(rsp)
}

func (c *ClientWithResponses) UpdateObjectWithResponse(ctx context.Context, objectType string, objectId string, params *UpdateObjectParams, body UpdateObjectJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateObjectResponse, error) {
	rsp, err := c.UpdateObject(ctx, objectType, objectId, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateObjectResponse(rsp)
}

// ParseGetObjectsResponse parses an HTTP response from a GetObjectsWithResponse call
func ParseGetObjectsResponse(rsp *http.Response) (*GetObjectsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetObjectsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateObjectResponse parses an HTTP response from a CreateObjectWithResponse call
func ParseCreateObjectResponse(rsp *http.Response) (*CreateObjectResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateObjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ObjectResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseSearchObjectsResponse parses an HTTP response from a SearchObjectsWithResponse call
func ParseSearchObjectsResponse(rsp *http.Response) (*SearchObjectsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SearchObjectsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ObjectsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteObjectResponse parses an HTTP response from a DeleteObjectWithResponse call
func ParseDeleteObjectResponse(rsp *http.Response) (*DeleteObjectResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteObjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetObjectByTypeAndIdResponse parses an HTTP response from a GetObjectByTypeAndIdWithResponse call
func ParseGetObjectByTypeAndIdResponse(rsp *http.Response) (*GetObjectByTypeAndIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetObjectByTypeAndIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ObjectResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateObjectResponse parses an HTTP response from a UpdateObjectWithResponse call
func ParseUpdateObjectResponse(rsp *http.Response) (*UpdateObjectResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateObjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the object was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the object was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the updated object.
			Id string `json:"id,omitempty"`

			// Properties Properties of the updated object.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the object's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the object was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
