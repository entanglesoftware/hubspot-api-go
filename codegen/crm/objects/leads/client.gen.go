// Package leads provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package leads

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
	// GetLeads request
	GetLeads(ctx context.Context, params *GetLeadsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateLeadWithBody request with any body
	CreateLeadWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateLead(ctx context.Context, body CreateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SearchLeadsWithBody request with any body
	SearchLeadsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SearchLeads(ctx context.Context, body SearchLeadsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteLeadById request
	DeleteLeadById(ctx context.Context, leadId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLeadById request
	GetLeadById(ctx context.Context, leadId string, params *GetLeadByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateLeadWithBody request with any body
	UpdateLeadWithBody(ctx context.Context, leadId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateLead(ctx context.Context, leadId string, body UpdateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetLeads(ctx context.Context, params *GetLeadsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLeadsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateLeadWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateLeadRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateLead(ctx context.Context, body CreateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateLeadRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchLeadsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchLeadsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchLeads(ctx context.Context, body SearchLeadsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchLeadsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteLeadById(ctx context.Context, leadId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteLeadByIdRequest(c.Server, leadId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetLeadById(ctx context.Context, leadId string, params *GetLeadByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLeadByIdRequest(c.Server, leadId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateLeadWithBody(ctx context.Context, leadId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateLeadRequestWithBody(c.Server, leadId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateLead(ctx context.Context, leadId string, body UpdateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateLeadRequest(c.Server, leadId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetLeadsRequest generates requests for GetLeads
func NewGetLeadsRequest(server string, params *GetLeadsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/leads")
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

// NewCreateLeadRequest calls the generic CreateLead builder with application/json body
func NewCreateLeadRequest(server string, body CreateLeadJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateLeadRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateLeadRequestWithBody generates requests for CreateLead with any type of body
func NewCreateLeadRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/leads")
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

// NewSearchLeadsRequest calls the generic SearchLeads builder with application/json body
func NewSearchLeadsRequest(server string, body SearchLeadsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSearchLeadsRequestWithBody(server, "application/json", bodyReader)
}

// NewSearchLeadsRequestWithBody generates requests for SearchLeads with any type of body
func NewSearchLeadsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/leads/search")
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

// NewDeleteLeadByIdRequest generates requests for DeleteLeadById
func NewDeleteLeadByIdRequest(server string, leadId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "leadId", runtime.ParamLocationPath, leadId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/leads/%s", pathParam0)
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

// NewGetLeadByIdRequest generates requests for GetLeadById
func NewGetLeadByIdRequest(server string, leadId string, params *GetLeadByIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "leadId", runtime.ParamLocationPath, leadId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/leads/%s", pathParam0)
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

// NewUpdateLeadRequest calls the generic UpdateLead builder with application/json body
func NewUpdateLeadRequest(server string, leadId string, body UpdateLeadJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateLeadRequestWithBody(server, leadId, "application/json", bodyReader)
}

// NewUpdateLeadRequestWithBody generates requests for UpdateLead with any type of body
func NewUpdateLeadRequestWithBody(server string, leadId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "leadId", runtime.ParamLocationPath, leadId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/leads/%s", pathParam0)
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
	// GetLeadsWithResponse request
	GetLeadsWithResponse(ctx context.Context, params *GetLeadsParams, reqEditors ...RequestEditorFn) (*GetLeadsResponse, error)

	// CreateLeadWithBodyWithResponse request with any body
	CreateLeadWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateLeadResponse, error)

	CreateLeadWithResponse(ctx context.Context, body CreateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateLeadResponse, error)

	// SearchLeadsWithBodyWithResponse request with any body
	SearchLeadsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchLeadsResponse, error)

	SearchLeadsWithResponse(ctx context.Context, body SearchLeadsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchLeadsResponse, error)

	// DeleteLeadByIdWithResponse request
	DeleteLeadByIdWithResponse(ctx context.Context, leadId string, reqEditors ...RequestEditorFn) (*DeleteLeadByIdResponse, error)

	// GetLeadByIdWithResponse request
	GetLeadByIdWithResponse(ctx context.Context, leadId string, params *GetLeadByIdParams, reqEditors ...RequestEditorFn) (*GetLeadByIdResponse, error)

	// UpdateLeadWithBodyWithResponse request with any body
	UpdateLeadWithBodyWithResponse(ctx context.Context, leadId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLeadResponse, error)

	UpdateLeadWithResponse(ctx context.Context, leadId string, body UpdateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLeadResponse, error)
}

type GetLeadsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LeadsResponse
}

// Status returns HTTPResponse.Status
func (r GetLeadsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLeadsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateLeadResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the lead was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the lead was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the created lead.
		Id string `json:"id,omitempty"`

		// Properties Properties of the created lead.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the lead's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the lead was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateLeadResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateLeadResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SearchLeadsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LeadsResponse
}

// Status returns HTTPResponse.Status
func (r SearchLeadsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchLeadsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteLeadByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteLeadByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteLeadByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetLeadByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LeadResponse
}

// Status returns HTTPResponse.Status
func (r GetLeadByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLeadByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateLeadResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the lead was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the lead was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the updated lead.
		Id string `json:"id,omitempty"`

		// Properties Properties of the updated lead.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the lead's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the lead was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateLeadResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateLeadResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetLeadsWithResponse request returning *GetLeadsResponse
func (c *ClientWithResponses) GetLeadsWithResponse(ctx context.Context, params *GetLeadsParams, reqEditors ...RequestEditorFn) (*GetLeadsResponse, error) {
	rsp, err := c.GetLeads(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLeadsResponse(rsp)
}

// CreateLeadWithBodyWithResponse request with arbitrary body returning *CreateLeadResponse
func (c *ClientWithResponses) CreateLeadWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateLeadResponse, error) {
	rsp, err := c.CreateLeadWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateLeadResponse(rsp)
}

func (c *ClientWithResponses) CreateLeadWithResponse(ctx context.Context, body CreateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateLeadResponse, error) {
	rsp, err := c.CreateLead(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateLeadResponse(rsp)
}

// SearchLeadsWithBodyWithResponse request with arbitrary body returning *SearchLeadsResponse
func (c *ClientWithResponses) SearchLeadsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchLeadsResponse, error) {
	rsp, err := c.SearchLeadsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchLeadsResponse(rsp)
}

func (c *ClientWithResponses) SearchLeadsWithResponse(ctx context.Context, body SearchLeadsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchLeadsResponse, error) {
	rsp, err := c.SearchLeads(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchLeadsResponse(rsp)
}

// DeleteLeadByIdWithResponse request returning *DeleteLeadByIdResponse
func (c *ClientWithResponses) DeleteLeadByIdWithResponse(ctx context.Context, leadId string, reqEditors ...RequestEditorFn) (*DeleteLeadByIdResponse, error) {
	rsp, err := c.DeleteLeadById(ctx, leadId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteLeadByIdResponse(rsp)
}

// GetLeadByIdWithResponse request returning *GetLeadByIdResponse
func (c *ClientWithResponses) GetLeadByIdWithResponse(ctx context.Context, leadId string, params *GetLeadByIdParams, reqEditors ...RequestEditorFn) (*GetLeadByIdResponse, error) {
	rsp, err := c.GetLeadById(ctx, leadId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLeadByIdResponse(rsp)
}

// UpdateLeadWithBodyWithResponse request with arbitrary body returning *UpdateLeadResponse
func (c *ClientWithResponses) UpdateLeadWithBodyWithResponse(ctx context.Context, leadId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLeadResponse, error) {
	rsp, err := c.UpdateLeadWithBody(ctx, leadId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateLeadResponse(rsp)
}

func (c *ClientWithResponses) UpdateLeadWithResponse(ctx context.Context, leadId string, body UpdateLeadJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLeadResponse, error) {
	rsp, err := c.UpdateLead(ctx, leadId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateLeadResponse(rsp)
}

// ParseGetLeadsResponse parses an HTTP response from a GetLeadsWithResponse call
func ParseGetLeadsResponse(rsp *http.Response) (*GetLeadsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLeadsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LeadsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateLeadResponse parses an HTTP response from a CreateLeadWithResponse call
func ParseCreateLeadResponse(rsp *http.Response) (*CreateLeadResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateLeadResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the lead was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the lead was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the created lead.
			Id string `json:"id,omitempty"`

			// Properties Properties of the created lead.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the lead's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the lead was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseSearchLeadsResponse parses an HTTP response from a SearchLeadsWithResponse call
func ParseSearchLeadsResponse(rsp *http.Response) (*SearchLeadsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SearchLeadsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LeadsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteLeadByIdResponse parses an HTTP response from a DeleteLeadByIdWithResponse call
func ParseDeleteLeadByIdResponse(rsp *http.Response) (*DeleteLeadByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteLeadByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetLeadByIdResponse parses an HTTP response from a GetLeadByIdWithResponse call
func ParseGetLeadByIdResponse(rsp *http.Response) (*GetLeadByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLeadByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LeadResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateLeadResponse parses an HTTP response from a UpdateLeadWithResponse call
func ParseUpdateLeadResponse(rsp *http.Response) (*UpdateLeadResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateLeadResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the lead was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the lead was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the updated lead.
			Id string `json:"id,omitempty"`

			// Properties Properties of the updated lead.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the lead's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the lead was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
