// Package schemas provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package schemas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

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
	// GetAssociationsSchema request
	GetAssociationsSchema(ctx context.Context, fromObjectType string, toObjectType string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateAssociationSchemaWithBody request with any body
	CreateAssociationSchemaWithBody(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAssociationSchema(ctx context.Context, fromObjectType string, toObjectType string, body CreateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateAssociationSchemaWithBody request with any body
	UpdateAssociationSchemaWithBody(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAssociationSchema(ctx context.Context, fromObjectType string, toObjectType string, body UpdateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteAssociationsSchema request
	DeleteAssociationsSchema(ctx context.Context, fromObjectType string, toObjectType string, associationTypeId int, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAssociationsSchema(ctx context.Context, fromObjectType string, toObjectType string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAssociationsSchemaRequest(c.Server, fromObjectType, toObjectType)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAssociationSchemaWithBody(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAssociationSchemaRequestWithBody(c.Server, fromObjectType, toObjectType, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAssociationSchema(ctx context.Context, fromObjectType string, toObjectType string, body CreateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAssociationSchemaRequest(c.Server, fromObjectType, toObjectType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAssociationSchemaWithBody(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAssociationSchemaRequestWithBody(c.Server, fromObjectType, toObjectType, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAssociationSchema(ctx context.Context, fromObjectType string, toObjectType string, body UpdateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAssociationSchemaRequest(c.Server, fromObjectType, toObjectType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteAssociationsSchema(ctx context.Context, fromObjectType string, toObjectType string, associationTypeId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteAssociationsSchemaRequest(c.Server, fromObjectType, toObjectType, associationTypeId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAssociationsSchemaRequest generates requests for GetAssociationsSchema
func NewGetAssociationsSchemaRequest(server string, fromObjectType string, toObjectType string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "fromObjectType", runtime.ParamLocationPath, fromObjectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "toObjectType", runtime.ParamLocationPath, toObjectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateAssociationSchemaRequest calls the generic CreateAssociationSchema builder with application/json body
func NewCreateAssociationSchemaRequest(server string, fromObjectType string, toObjectType string, body CreateAssociationSchemaJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateAssociationSchemaRequestWithBody(server, fromObjectType, toObjectType, "application/json", bodyReader)
}

// NewCreateAssociationSchemaRequestWithBody generates requests for CreateAssociationSchema with any type of body
func NewCreateAssociationSchemaRequestWithBody(server string, fromObjectType string, toObjectType string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "fromObjectType", runtime.ParamLocationPath, fromObjectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "toObjectType", runtime.ParamLocationPath, toObjectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", pathParam0, pathParam1)
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

// NewUpdateAssociationSchemaRequest calls the generic UpdateAssociationSchema builder with application/json body
func NewUpdateAssociationSchemaRequest(server string, fromObjectType string, toObjectType string, body UpdateAssociationSchemaJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateAssociationSchemaRequestWithBody(server, fromObjectType, toObjectType, "application/json", bodyReader)
}

// NewUpdateAssociationSchemaRequestWithBody generates requests for UpdateAssociationSchema with any type of body
func NewUpdateAssociationSchemaRequestWithBody(server string, fromObjectType string, toObjectType string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "fromObjectType", runtime.ParamLocationPath, fromObjectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "toObjectType", runtime.ParamLocationPath, toObjectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteAssociationsSchemaRequest generates requests for DeleteAssociationsSchema
func NewDeleteAssociationsSchemaRequest(server string, fromObjectType string, toObjectType string, associationTypeId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "fromObjectType", runtime.ParamLocationPath, fromObjectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "toObjectType", runtime.ParamLocationPath, toObjectType)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "associationTypeId", runtime.ParamLocationPath, associationTypeId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v4/associations/%s/%s/labels/%s", pathParam0, pathParam1, pathParam2)
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
	// GetAssociationsSchemaWithResponse request
	GetAssociationsSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, reqEditors ...RequestEditorFn) (*GetAssociationsSchemaResponse, error)

	// CreateAssociationSchemaWithBodyWithResponse request with any body
	CreateAssociationSchemaWithBodyWithResponse(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAssociationSchemaResponse, error)

	CreateAssociationSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, body CreateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAssociationSchemaResponse, error)

	// UpdateAssociationSchemaWithBodyWithResponse request with any body
	UpdateAssociationSchemaWithBodyWithResponse(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAssociationSchemaResponse, error)

	UpdateAssociationSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, body UpdateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAssociationSchemaResponse, error)

	// DeleteAssociationsSchemaWithResponse request
	DeleteAssociationsSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, associationTypeId int, reqEditors ...RequestEditorFn) (*DeleteAssociationsSchemaResponse, error)
}

type GetAssociationsSchemaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AssociationsResponse
}

// Status returns HTTPResponse.Status
func (r GetAssociationsSchemaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAssociationsSchemaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateAssociationSchemaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AssociationsResponse
}

// Status returns HTTPResponse.Status
func (r CreateAssociationSchemaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAssociationSchemaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateAssociationSchemaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r UpdateAssociationSchemaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateAssociationSchemaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteAssociationsSchemaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteAssociationsSchemaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteAssociationsSchemaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAssociationsSchemaWithResponse request returning *GetAssociationsSchemaResponse
func (c *ClientWithResponses) GetAssociationsSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, reqEditors ...RequestEditorFn) (*GetAssociationsSchemaResponse, error) {
	rsp, err := c.GetAssociationsSchema(ctx, fromObjectType, toObjectType, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAssociationsSchemaResponse(rsp)
}

// CreateAssociationSchemaWithBodyWithResponse request with arbitrary body returning *CreateAssociationSchemaResponse
func (c *ClientWithResponses) CreateAssociationSchemaWithBodyWithResponse(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAssociationSchemaResponse, error) {
	rsp, err := c.CreateAssociationSchemaWithBody(ctx, fromObjectType, toObjectType, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAssociationSchemaResponse(rsp)
}

func (c *ClientWithResponses) CreateAssociationSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, body CreateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAssociationSchemaResponse, error) {
	rsp, err := c.CreateAssociationSchema(ctx, fromObjectType, toObjectType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAssociationSchemaResponse(rsp)
}

// UpdateAssociationSchemaWithBodyWithResponse request with arbitrary body returning *UpdateAssociationSchemaResponse
func (c *ClientWithResponses) UpdateAssociationSchemaWithBodyWithResponse(ctx context.Context, fromObjectType string, toObjectType string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAssociationSchemaResponse, error) {
	rsp, err := c.UpdateAssociationSchemaWithBody(ctx, fromObjectType, toObjectType, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAssociationSchemaResponse(rsp)
}

func (c *ClientWithResponses) UpdateAssociationSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, body UpdateAssociationSchemaJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAssociationSchemaResponse, error) {
	rsp, err := c.UpdateAssociationSchema(ctx, fromObjectType, toObjectType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAssociationSchemaResponse(rsp)
}

// DeleteAssociationsSchemaWithResponse request returning *DeleteAssociationsSchemaResponse
func (c *ClientWithResponses) DeleteAssociationsSchemaWithResponse(ctx context.Context, fromObjectType string, toObjectType string, associationTypeId int, reqEditors ...RequestEditorFn) (*DeleteAssociationsSchemaResponse, error) {
	rsp, err := c.DeleteAssociationsSchema(ctx, fromObjectType, toObjectType, associationTypeId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteAssociationsSchemaResponse(rsp)
}

// ParseGetAssociationsSchemaResponse parses an HTTP response from a GetAssociationsSchemaWithResponse call
func ParseGetAssociationsSchemaResponse(rsp *http.Response) (*GetAssociationsSchemaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAssociationsSchemaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AssociationsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateAssociationSchemaResponse parses an HTTP response from a CreateAssociationSchemaWithResponse call
func ParseCreateAssociationSchemaResponse(rsp *http.Response) (*CreateAssociationSchemaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAssociationSchemaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AssociationsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateAssociationSchemaResponse parses an HTTP response from a UpdateAssociationSchemaWithResponse call
func ParseUpdateAssociationSchemaResponse(rsp *http.Response) (*UpdateAssociationSchemaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateAssociationSchemaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseDeleteAssociationsSchemaResponse parses an HTTP response from a DeleteAssociationsSchemaWithResponse call
func ParseDeleteAssociationsSchemaResponse(rsp *http.Response) (*DeleteAssociationsSchemaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteAssociationsSchemaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
