// Package tickets provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package tickets

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
	// GetTickets request
	GetTickets(ctx context.Context, params *GetTicketsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateTicketWithBody request with any body
	CreateTicketWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateTicket(ctx context.Context, body CreateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SearchTicketsWithBody request with any body
	SearchTicketsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SearchTickets(ctx context.Context, body SearchTicketsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteTicketById request
	DeleteTicketById(ctx context.Context, ticketId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTicketById request
	GetTicketById(ctx context.Context, ticketId string, params *GetTicketByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateTicketWithBody request with any body
	UpdateTicketWithBody(ctx context.Context, ticketId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateTicket(ctx context.Context, ticketId string, body UpdateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetTickets(ctx context.Context, params *GetTicketsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTicketsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateTicketWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTicketRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateTicket(ctx context.Context, body CreateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTicketRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchTicketsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchTicketsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchTickets(ctx context.Context, body SearchTicketsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchTicketsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteTicketById(ctx context.Context, ticketId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteTicketByIdRequest(c.Server, ticketId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTicketById(ctx context.Context, ticketId string, params *GetTicketByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTicketByIdRequest(c.Server, ticketId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateTicketWithBody(ctx context.Context, ticketId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateTicketRequestWithBody(c.Server, ticketId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateTicket(ctx context.Context, ticketId string, body UpdateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateTicketRequest(c.Server, ticketId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetTicketsRequest generates requests for GetTickets
func NewGetTicketsRequest(server string, params *GetTicketsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/tickets")
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

// NewCreateTicketRequest calls the generic CreateTicket builder with application/json body
func NewCreateTicketRequest(server string, body CreateTicketJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateTicketRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateTicketRequestWithBody generates requests for CreateTicket with any type of body
func NewCreateTicketRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/tickets")
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

// NewSearchTicketsRequest calls the generic SearchTickets builder with application/json body
func NewSearchTicketsRequest(server string, body SearchTicketsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSearchTicketsRequestWithBody(server, "application/json", bodyReader)
}

// NewSearchTicketsRequestWithBody generates requests for SearchTickets with any type of body
func NewSearchTicketsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/tickets/search")
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

// NewDeleteTicketByIdRequest generates requests for DeleteTicketById
func NewDeleteTicketByIdRequest(server string, ticketId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ticketId", runtime.ParamLocationPath, ticketId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/tickets/%s", pathParam0)
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

// NewGetTicketByIdRequest generates requests for GetTicketById
func NewGetTicketByIdRequest(server string, ticketId string, params *GetTicketByIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ticketId", runtime.ParamLocationPath, ticketId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/tickets/%s", pathParam0)
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

// NewUpdateTicketRequest calls the generic UpdateTicket builder with application/json body
func NewUpdateTicketRequest(server string, ticketId string, body UpdateTicketJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateTicketRequestWithBody(server, ticketId, "application/json", bodyReader)
}

// NewUpdateTicketRequestWithBody generates requests for UpdateTicket with any type of body
func NewUpdateTicketRequestWithBody(server string, ticketId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ticketId", runtime.ParamLocationPath, ticketId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/tickets/%s", pathParam0)
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
	// GetTicketsWithResponse request
	GetTicketsWithResponse(ctx context.Context, params *GetTicketsParams, reqEditors ...RequestEditorFn) (*GetTicketsResponse, error)

	// CreateTicketWithBodyWithResponse request with any body
	CreateTicketWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTicketResponse, error)

	CreateTicketWithResponse(ctx context.Context, body CreateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTicketResponse, error)

	// SearchTicketsWithBodyWithResponse request with any body
	SearchTicketsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchTicketsResponse, error)

	SearchTicketsWithResponse(ctx context.Context, body SearchTicketsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchTicketsResponse, error)

	// DeleteTicketByIdWithResponse request
	DeleteTicketByIdWithResponse(ctx context.Context, ticketId string, reqEditors ...RequestEditorFn) (*DeleteTicketByIdResponse, error)

	// GetTicketByIdWithResponse request
	GetTicketByIdWithResponse(ctx context.Context, ticketId string, params *GetTicketByIdParams, reqEditors ...RequestEditorFn) (*GetTicketByIdResponse, error)

	// UpdateTicketWithBodyWithResponse request with any body
	UpdateTicketWithBodyWithResponse(ctx context.Context, ticketId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTicketResponse, error)

	UpdateTicketWithResponse(ctx context.Context, ticketId string, body UpdateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTicketResponse, error)
}

type GetTicketsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *TicketsResponse
}

// Status returns HTTPResponse.Status
func (r GetTicketsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTicketsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateTicketResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the ticket was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the ticket was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the created ticket.
		Id string `json:"id,omitempty"`

		// Properties Properties of the created ticket.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the ticket's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the ticket was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateTicketResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateTicketResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SearchTicketsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *TicketsResponse
}

// Status returns HTTPResponse.Status
func (r SearchTicketsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchTicketsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteTicketByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteTicketByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteTicketByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTicketByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *TicketResponse
}

// Status returns HTTPResponse.Status
func (r GetTicketByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTicketByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateTicketResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the ticket was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the ticket was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the updated ticket.
		Id string `json:"id,omitempty"`

		// Properties Properties of the updated ticket.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the ticket's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the ticket was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateTicketResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateTicketResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetTicketsWithResponse request returning *GetTicketsResponse
func (c *ClientWithResponses) GetTicketsWithResponse(ctx context.Context, params *GetTicketsParams, reqEditors ...RequestEditorFn) (*GetTicketsResponse, error) {
	rsp, err := c.GetTickets(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTicketsResponse(rsp)
}

// CreateTicketWithBodyWithResponse request with arbitrary body returning *CreateTicketResponse
func (c *ClientWithResponses) CreateTicketWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTicketResponse, error) {
	rsp, err := c.CreateTicketWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTicketResponse(rsp)
}

func (c *ClientWithResponses) CreateTicketWithResponse(ctx context.Context, body CreateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTicketResponse, error) {
	rsp, err := c.CreateTicket(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTicketResponse(rsp)
}

// SearchTicketsWithBodyWithResponse request with arbitrary body returning *SearchTicketsResponse
func (c *ClientWithResponses) SearchTicketsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchTicketsResponse, error) {
	rsp, err := c.SearchTicketsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchTicketsResponse(rsp)
}

func (c *ClientWithResponses) SearchTicketsWithResponse(ctx context.Context, body SearchTicketsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchTicketsResponse, error) {
	rsp, err := c.SearchTickets(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchTicketsResponse(rsp)
}

// DeleteTicketByIdWithResponse request returning *DeleteTicketByIdResponse
func (c *ClientWithResponses) DeleteTicketByIdWithResponse(ctx context.Context, ticketId string, reqEditors ...RequestEditorFn) (*DeleteTicketByIdResponse, error) {
	rsp, err := c.DeleteTicketById(ctx, ticketId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteTicketByIdResponse(rsp)
}

// GetTicketByIdWithResponse request returning *GetTicketByIdResponse
func (c *ClientWithResponses) GetTicketByIdWithResponse(ctx context.Context, ticketId string, params *GetTicketByIdParams, reqEditors ...RequestEditorFn) (*GetTicketByIdResponse, error) {
	rsp, err := c.GetTicketById(ctx, ticketId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTicketByIdResponse(rsp)
}

// UpdateTicketWithBodyWithResponse request with arbitrary body returning *UpdateTicketResponse
func (c *ClientWithResponses) UpdateTicketWithBodyWithResponse(ctx context.Context, ticketId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTicketResponse, error) {
	rsp, err := c.UpdateTicketWithBody(ctx, ticketId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateTicketResponse(rsp)
}

func (c *ClientWithResponses) UpdateTicketWithResponse(ctx context.Context, ticketId string, body UpdateTicketJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTicketResponse, error) {
	rsp, err := c.UpdateTicket(ctx, ticketId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateTicketResponse(rsp)
}

// ParseGetTicketsResponse parses an HTTP response from a GetTicketsWithResponse call
func ParseGetTicketsResponse(rsp *http.Response) (*GetTicketsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTicketsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest TicketsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateTicketResponse parses an HTTP response from a CreateTicketWithResponse call
func ParseCreateTicketResponse(rsp *http.Response) (*CreateTicketResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateTicketResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the ticket was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the ticket was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the created ticket.
			Id string `json:"id,omitempty"`

			// Properties Properties of the created ticket.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the ticket's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the ticket was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseSearchTicketsResponse parses an HTTP response from a SearchTicketsWithResponse call
func ParseSearchTicketsResponse(rsp *http.Response) (*SearchTicketsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SearchTicketsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest TicketsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteTicketByIdResponse parses an HTTP response from a DeleteTicketByIdWithResponse call
func ParseDeleteTicketByIdResponse(rsp *http.Response) (*DeleteTicketByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteTicketByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetTicketByIdResponse parses an HTTP response from a GetTicketByIdWithResponse call
func ParseGetTicketByIdResponse(rsp *http.Response) (*GetTicketByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTicketByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest TicketResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateTicketResponse parses an HTTP response from a UpdateTicketWithResponse call
func ParseUpdateTicketResponse(rsp *http.Response) (*UpdateTicketResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateTicketResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the ticket was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the ticket was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the updated ticket.
			Id string `json:"id,omitempty"`

			// Properties Properties of the updated ticket.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the ticket's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the ticket was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
