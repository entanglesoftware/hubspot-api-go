// Package products provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package products

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
	// GetProducts request
	GetProducts(ctx context.Context, params *GetProductsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateProductWithBody request with any body
	CreateProductWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateProduct(ctx context.Context, body CreateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BatchUpsertProductsWithBody request with any body
	BatchUpsertProductsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BatchUpsertProducts(ctx context.Context, body BatchUpsertProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SearchProductsWithBody request with any body
	SearchProductsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SearchProducts(ctx context.Context, body SearchProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteProductById request
	DeleteProductById(ctx context.Context, productId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetProductById request
	GetProductById(ctx context.Context, productId string, params *GetProductByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateProductWithBody request with any body
	UpdateProductWithBody(ctx context.Context, productId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateProduct(ctx context.Context, productId string, body UpdateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetProducts(ctx context.Context, params *GetProductsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProductsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateProductWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateProductRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateProduct(ctx context.Context, body CreateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateProductRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BatchUpsertProductsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBatchUpsertProductsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BatchUpsertProducts(ctx context.Context, body BatchUpsertProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBatchUpsertProductsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchProductsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchProductsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchProducts(ctx context.Context, body SearchProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchProductsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteProductById(ctx context.Context, productId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteProductByIdRequest(c.Server, productId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetProductById(ctx context.Context, productId string, params *GetProductByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProductByIdRequest(c.Server, productId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateProductWithBody(ctx context.Context, productId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateProductRequestWithBody(c.Server, productId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateProduct(ctx context.Context, productId string, body UpdateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateProductRequest(c.Server, productId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetProductsRequest generates requests for GetProducts
func NewGetProductsRequest(server string, params *GetProductsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/products")
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

// NewCreateProductRequest calls the generic CreateProduct builder with application/json body
func NewCreateProductRequest(server string, body CreateProductJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateProductRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateProductRequestWithBody generates requests for CreateProduct with any type of body
func NewCreateProductRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/products")
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

// NewBatchUpsertProductsRequest calls the generic BatchUpsertProducts builder with application/json body
func NewBatchUpsertProductsRequest(server string, body BatchUpsertProductsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBatchUpsertProductsRequestWithBody(server, "application/json", bodyReader)
}

// NewBatchUpsertProductsRequestWithBody generates requests for BatchUpsertProducts with any type of body
func NewBatchUpsertProductsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/products/batch/upsert")
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

// NewSearchProductsRequest calls the generic SearchProducts builder with application/json body
func NewSearchProductsRequest(server string, body SearchProductsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSearchProductsRequestWithBody(server, "application/json", bodyReader)
}

// NewSearchProductsRequestWithBody generates requests for SearchProducts with any type of body
func NewSearchProductsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/products/search")
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

// NewDeleteProductByIdRequest generates requests for DeleteProductById
func NewDeleteProductByIdRequest(server string, productId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "productId", runtime.ParamLocationPath, productId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/products/%s", pathParam0)
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

// NewGetProductByIdRequest generates requests for GetProductById
func NewGetProductByIdRequest(server string, productId string, params *GetProductByIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "productId", runtime.ParamLocationPath, productId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/products/%s", pathParam0)
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

// NewUpdateProductRequest calls the generic UpdateProduct builder with application/json body
func NewUpdateProductRequest(server string, productId string, body UpdateProductJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateProductRequestWithBody(server, productId, "application/json", bodyReader)
}

// NewUpdateProductRequestWithBody generates requests for UpdateProduct with any type of body
func NewUpdateProductRequestWithBody(server string, productId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "productId", runtime.ParamLocationPath, productId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/products/%s", pathParam0)
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
	// GetProductsWithResponse request
	GetProductsWithResponse(ctx context.Context, params *GetProductsParams, reqEditors ...RequestEditorFn) (*GetProductsResponse, error)

	// CreateProductWithBodyWithResponse request with any body
	CreateProductWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateProductResponse, error)

	CreateProductWithResponse(ctx context.Context, body CreateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateProductResponse, error)

	// BatchUpsertProductsWithBodyWithResponse request with any body
	BatchUpsertProductsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BatchUpsertProductsResponse, error)

	BatchUpsertProductsWithResponse(ctx context.Context, body BatchUpsertProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*BatchUpsertProductsResponse, error)

	// SearchProductsWithBodyWithResponse request with any body
	SearchProductsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchProductsResponse, error)

	SearchProductsWithResponse(ctx context.Context, body SearchProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchProductsResponse, error)

	// DeleteProductByIdWithResponse request
	DeleteProductByIdWithResponse(ctx context.Context, productId string, reqEditors ...RequestEditorFn) (*DeleteProductByIdResponse, error)

	// GetProductByIdWithResponse request
	GetProductByIdWithResponse(ctx context.Context, productId string, params *GetProductByIdParams, reqEditors ...RequestEditorFn) (*GetProductByIdResponse, error)

	// UpdateProductWithBodyWithResponse request with any body
	UpdateProductWithBodyWithResponse(ctx context.Context, productId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateProductResponse, error)

	UpdateProductWithResponse(ctx context.Context, productId string, body UpdateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateProductResponse, error)
}

type GetProductsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ProductsResponse
}

// Status returns HTTPResponse.Status
func (r GetProductsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProductsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateProductResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the product was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the product was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the created product.
		Id string `json:"id,omitempty"`

		// Properties Properties of the created product.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the product's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the product was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateProductResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateProductResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BatchUpsertProductsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BatchProductsResponse
	JSON400      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r BatchUpsertProductsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BatchUpsertProductsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SearchProductsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ProductsResponse
}

// Status returns HTTPResponse.Status
func (r SearchProductsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchProductsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteProductByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteProductByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteProductByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProductByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ProductResponse
}

// Status returns HTTPResponse.Status
func (r GetProductByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProductByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateProductResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Archived Whether the customer is archived or not.
		Archived bool `json:"archived,omitempty"`

		// ArchivedAt Timestamp when the product was archived.
		ArchivedAt time.Time `json:"archivedAt,omitempty"`

		// CreatedAt Timestamp when the product was created.
		CreatedAt time.Time `json:"createdAt,omitempty"`

		// Id Unique ID of the updated product.
		Id string `json:"id,omitempty"`

		// Properties Properties of the updated product.
		Properties map[string]interface{} `json:"properties,omitempty"`

		// PropertiesWithHistory A map of the product's properties including historical values.
		PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

		// UpdatedAt Timestamp when the product was last updated.
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateProductResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateProductResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetProductsWithResponse request returning *GetProductsResponse
func (c *ClientWithResponses) GetProductsWithResponse(ctx context.Context, params *GetProductsParams, reqEditors ...RequestEditorFn) (*GetProductsResponse, error) {
	rsp, err := c.GetProducts(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProductsResponse(rsp)
}

// CreateProductWithBodyWithResponse request with arbitrary body returning *CreateProductResponse
func (c *ClientWithResponses) CreateProductWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateProductResponse, error) {
	rsp, err := c.CreateProductWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateProductResponse(rsp)
}

func (c *ClientWithResponses) CreateProductWithResponse(ctx context.Context, body CreateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateProductResponse, error) {
	rsp, err := c.CreateProduct(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateProductResponse(rsp)
}

// BatchUpsertProductsWithBodyWithResponse request with arbitrary body returning *BatchUpsertProductsResponse
func (c *ClientWithResponses) BatchUpsertProductsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BatchUpsertProductsResponse, error) {
	rsp, err := c.BatchUpsertProductsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBatchUpsertProductsResponse(rsp)
}

func (c *ClientWithResponses) BatchUpsertProductsWithResponse(ctx context.Context, body BatchUpsertProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*BatchUpsertProductsResponse, error) {
	rsp, err := c.BatchUpsertProducts(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBatchUpsertProductsResponse(rsp)
}

// SearchProductsWithBodyWithResponse request with arbitrary body returning *SearchProductsResponse
func (c *ClientWithResponses) SearchProductsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchProductsResponse, error) {
	rsp, err := c.SearchProductsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchProductsResponse(rsp)
}

func (c *ClientWithResponses) SearchProductsWithResponse(ctx context.Context, body SearchProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchProductsResponse, error) {
	rsp, err := c.SearchProducts(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchProductsResponse(rsp)
}

// DeleteProductByIdWithResponse request returning *DeleteProductByIdResponse
func (c *ClientWithResponses) DeleteProductByIdWithResponse(ctx context.Context, productId string, reqEditors ...RequestEditorFn) (*DeleteProductByIdResponse, error) {
	rsp, err := c.DeleteProductById(ctx, productId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteProductByIdResponse(rsp)
}

// GetProductByIdWithResponse request returning *GetProductByIdResponse
func (c *ClientWithResponses) GetProductByIdWithResponse(ctx context.Context, productId string, params *GetProductByIdParams, reqEditors ...RequestEditorFn) (*GetProductByIdResponse, error) {
	rsp, err := c.GetProductById(ctx, productId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProductByIdResponse(rsp)
}

// UpdateProductWithBodyWithResponse request with arbitrary body returning *UpdateProductResponse
func (c *ClientWithResponses) UpdateProductWithBodyWithResponse(ctx context.Context, productId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateProductResponse, error) {
	rsp, err := c.UpdateProductWithBody(ctx, productId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateProductResponse(rsp)
}

func (c *ClientWithResponses) UpdateProductWithResponse(ctx context.Context, productId string, body UpdateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateProductResponse, error) {
	rsp, err := c.UpdateProduct(ctx, productId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateProductResponse(rsp)
}

// ParseGetProductsResponse parses an HTTP response from a GetProductsWithResponse call
func ParseGetProductsResponse(rsp *http.Response) (*GetProductsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetProductsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ProductsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateProductResponse parses an HTTP response from a CreateProductWithResponse call
func ParseCreateProductResponse(rsp *http.Response) (*CreateProductResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateProductResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the product was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the product was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the created product.
			Id string `json:"id,omitempty"`

			// Properties Properties of the created product.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the product's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the product was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseBatchUpsertProductsResponse parses an HTTP response from a BatchUpsertProductsWithResponse call
func ParseBatchUpsertProductsResponse(rsp *http.Response) (*BatchUpsertProductsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BatchUpsertProductsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BatchProductsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseSearchProductsResponse parses an HTTP response from a SearchProductsWithResponse call
func ParseSearchProductsResponse(rsp *http.Response) (*SearchProductsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SearchProductsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ProductsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteProductByIdResponse parses an HTTP response from a DeleteProductByIdWithResponse call
func ParseDeleteProductByIdResponse(rsp *http.Response) (*DeleteProductByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteProductByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetProductByIdResponse parses an HTTP response from a GetProductByIdWithResponse call
func ParseGetProductByIdResponse(rsp *http.Response) (*GetProductByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetProductByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ProductResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateProductResponse parses an HTTP response from a UpdateProductWithResponse call
func ParseUpdateProductResponse(rsp *http.Response) (*UpdateProductResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateProductResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Archived Whether the customer is archived or not.
			Archived bool `json:"archived,omitempty"`

			// ArchivedAt Timestamp when the product was archived.
			ArchivedAt time.Time `json:"archivedAt,omitempty"`

			// CreatedAt Timestamp when the product was created.
			CreatedAt time.Time `json:"createdAt,omitempty"`

			// Id Unique ID of the updated product.
			Id string `json:"id,omitempty"`

			// Properties Properties of the updated product.
			Properties map[string]interface{} `json:"properties,omitempty"`

			// PropertiesWithHistory A map of the product's properties including historical values.
			PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

			// UpdatedAt Timestamp when the product was last updated.
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
