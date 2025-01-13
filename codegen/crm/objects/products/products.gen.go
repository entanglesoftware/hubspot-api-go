// Package products provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package products

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	Oauth2Scopes = "oauth2.Scopes"
)

// Defines values for CreateProductJSONBodyAssociationsTypesAssociationCategory.
const (
	HUBSPOTDEFINED    CreateProductJSONBodyAssociationsTypesAssociationCategory = "HUBSPOT_DEFINED"
	INTEGRATORDEFINED CreateProductJSONBodyAssociationsTypesAssociationCategory = "INTEGRATOR_DEFINED"
	Search            CreateProductJSONBodyAssociationsTypesAssociationCategory = "Search"
	USERDEFINED       CreateProductJSONBodyAssociationsTypesAssociationCategory = "USER_DEFINED"
)

// GetProductsParams defines parameters for GetProducts.
type GetProductsParams struct {
	// Limit Maximum number of results per page.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// After Cursor token for the next page of results.
	After *string `form:"after,omitempty" json:"after,omitempty"`

	// Properties Comma-separated list of properties to include in the response.
	// If a specified property is not present, it will be ignored.
	Properties *Properties `form:"properties,omitempty" json:"properties,omitempty"`

	// PropertiesWithHistory Comma-separated list of properties to include with their historical values.
	// Historical data reduces the maximum number of objects returned per request.
	PropertiesWithHistory *PropertiesWithHistory `form:"propertiesWithHistory,omitempty" json:"propertiesWithHistory,omitempty"`

	// Associations Comma-separated list of object types to retrieve associated IDs for.
	// Nonexistent associations will be ignored.
	Associations *Associations `form:"associations,omitempty" json:"associations,omitempty"`

	// Archived Include only archived results.
	Archived *Archived `form:"archived,omitempty" json:"archived,omitempty"`
}

// CreateProductJSONBody defines parameters for CreateProduct.
type CreateProductJSONBody struct {
	// Associations List of associations for the product.
	Associations *[]struct {
		// To Target object details for the association.
		To *struct {
			// Id Target object ID.
			Id *string `json:"id,omitempty"`
		} `json:"to,omitempty"`
		Types *[]struct {
			// AssociationCategory Category of the association.
			AssociationCategory *CreateProductJSONBodyAssociationsTypesAssociationCategory `json:"associationCategory,omitempty"`

			// AssociationTypeId ID of the association type.
			AssociationTypeId *int32 `json:"associationTypeId,omitempty"`
		} `json:"types,omitempty"`
	} `json:"associations,omitempty"`

	// ObjectWriteTraceId Trace ID for object write operations.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs of product properties.
	Properties map[string]string `json:"properties"`
}

// CreateProductJSONBodyAssociationsTypesAssociationCategory defines parameters for CreateProduct.
type CreateProductJSONBodyAssociationsTypesAssociationCategory string

// SearchProductsJSONBody defines parameters for SearchProducts.
type SearchProductsJSONBody struct {
	After      *string   `json:"after,omitempty"`
	Limit      *int      `json:"limit,omitempty"`
	Properties *[]string `json:"properties,omitempty"`
	Query      *string   `json:"query,omitempty"`
	Schema     *Filters  `json:"schema,omitempty"`
	Sorts      *[]string `json:"sorts,omitempty"`
}

// SearchProductsParams defines parameters for SearchProducts.
type SearchProductsParams struct {
	// Hapikey HubSpot API key
	Hapikey string `form:"hapikey" json:"hapikey"`
}

// GetProductByIdParams defines parameters for GetProductById.
type GetProductByIdParams struct {
	// IdProperty The property to use as the ID.
	IdProperty *string `form:"idProperty,omitempty" json:"idProperty,omitempty"`

	// Properties Comma-separated list of properties to include in the response.
	// If a specified property is not present, it will be ignored.
	Properties *Properties `form:"properties,omitempty" json:"properties,omitempty"`

	// PropertiesWithHistory Comma-separated list of properties to include with their historical values.
	// Historical data reduces the maximum number of objects returned per request.
	PropertiesWithHistory *PropertiesWithHistory `form:"propertiesWithHistory,omitempty" json:"propertiesWithHistory,omitempty"`

	// Associations Comma-separated list of object types to retrieve associated IDs for.
	// Nonexistent associations will be ignored.
	Associations *Associations `form:"associations,omitempty" json:"associations,omitempty"`

	// Archived Include only archived results.
	Archived *Archived `form:"archived,omitempty" json:"archived,omitempty"`
}

// UpdateProductJSONBody defines parameters for UpdateProduct.
type UpdateProductJSONBody struct {
	// ObjectWriteTraceId Unique trace ID for the operation.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs representing the deal properties to update.
	Properties map[string]string `json:"properties"`
}

// CreateProductJSONRequestBody defines body for CreateProduct for application/json ContentType.
type CreateProductJSONRequestBody CreateProductJSONBody

// SearchProductsJSONRequestBody defines body for SearchProducts for application/json ContentType.
type SearchProductsJSONRequestBody SearchProductsJSONBody

// UpdateProductJSONRequestBody defines body for UpdateProduct for application/json ContentType.
type UpdateProductJSONRequestBody UpdateProductJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of products
	// (GET /crm/v3/objects/products)
	GetProducts(ctx echo.Context, params GetProductsParams) error
	// Create a new product
	// (POST /crm/v3/objects/products)
	CreateProduct(ctx echo.Context) error
	// Search for products by email
	// (POST /crm/v3/objects/products/search)
	SearchProducts(ctx echo.Context, params SearchProductsParams) error
	// Delete a product
	// (DELETE /crm/v3/objects/products/{productId})
	DeleteProductById(ctx echo.Context, productId string) error
	// Get Product Details
	// (GET /crm/v3/objects/products/{productId})
	GetProductById(ctx echo.Context, productId string, params GetProductByIdParams) error
	// Update a product
	// (PATCH /crm/v3/objects/products/{productId})
	UpdateProduct(ctx echo.Context, productId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetProducts converts echo context to params.
func (w *ServerInterfaceWrapper) GetProducts(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProductsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", ctx.QueryParams(), &params.After)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter after: %s", err))
	}

	// ------------- Optional query parameter "properties" -------------

	err = runtime.BindQueryParameter("form", false, false, "properties", ctx.QueryParams(), &params.Properties)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter properties: %s", err))
	}

	// ------------- Optional query parameter "propertiesWithHistory" -------------

	err = runtime.BindQueryParameter("form", false, false, "propertiesWithHistory", ctx.QueryParams(), &params.PropertiesWithHistory)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter propertiesWithHistory: %s", err))
	}

	// ------------- Optional query parameter "associations" -------------

	err = runtime.BindQueryParameter("form", false, false, "associations", ctx.QueryParams(), &params.Associations)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter associations: %s", err))
	}

	// ------------- Optional query parameter "archived" -------------

	err = runtime.BindQueryParameter("form", true, false, "archived", ctx.QueryParams(), &params.Archived)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter archived: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProducts(ctx, params)
	return err
}

// CreateProduct converts echo context to params.
func (w *ServerInterfaceWrapper) CreateProduct(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateProduct(ctx)
	return err
}

// SearchProducts converts echo context to params.
func (w *ServerInterfaceWrapper) SearchProducts(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchProductsParams
	// ------------- Required query parameter "hapikey" -------------

	err = runtime.BindQueryParameter("form", true, true, "hapikey", ctx.QueryParams(), &params.Hapikey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hapikey: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchProducts(ctx, params)
	return err
}

// DeleteProductById converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteProductById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "productId" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "productId", ctx.Param("productId"), &productId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter productId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteProductById(ctx, productId)
	return err
}

// GetProductById converts echo context to params.
func (w *ServerInterfaceWrapper) GetProductById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "productId" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "productId", ctx.Param("productId"), &productId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter productId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProductByIdParams
	// ------------- Optional query parameter "idProperty" -------------

	err = runtime.BindQueryParameter("form", true, false, "idProperty", ctx.QueryParams(), &params.IdProperty)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter idProperty: %s", err))
	}

	// ------------- Optional query parameter "properties" -------------

	err = runtime.BindQueryParameter("form", false, false, "properties", ctx.QueryParams(), &params.Properties)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter properties: %s", err))
	}

	// ------------- Optional query parameter "propertiesWithHistory" -------------

	err = runtime.BindQueryParameter("form", false, false, "propertiesWithHistory", ctx.QueryParams(), &params.PropertiesWithHistory)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter propertiesWithHistory: %s", err))
	}

	// ------------- Optional query parameter "associations" -------------

	err = runtime.BindQueryParameter("form", false, false, "associations", ctx.QueryParams(), &params.Associations)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter associations: %s", err))
	}

	// ------------- Optional query parameter "archived" -------------

	err = runtime.BindQueryParameter("form", true, false, "archived", ctx.QueryParams(), &params.Archived)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter archived: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProductById(ctx, productId, params)
	return err
}

// UpdateProduct converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateProduct(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "productId" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "productId", ctx.Param("productId"), &productId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter productId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateProduct(ctx, productId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/crm/v3/objects/products", wrapper.GetProducts)
	router.POST(baseURL+"/crm/v3/objects/products", wrapper.CreateProduct)
	router.POST(baseURL+"/crm/v3/objects/products/search", wrapper.SearchProducts)
	router.DELETE(baseURL+"/crm/v3/objects/products/:productId", wrapper.DeleteProductById)
	router.GET(baseURL+"/crm/v3/objects/products/:productId", wrapper.GetProductById)
	router.PATCH(baseURL+"/crm/v3/objects/products/:productId", wrapper.UpdateProduct)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{
	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
