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
	After        *string       `json:"after,omitempty"`
	FilterGroups *FilterGroups `json:"filterGroups,omitempty"`
	Limit        *int          `json:"limit,omitempty"`
	Properties   *[]string     `json:"properties,omitempty"`
	Query        *string       `json:"query,omitempty"`
	Sorts        *[]string     `json:"sorts,omitempty"`
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
	"6r8BAAD//2psukQTNAAA",
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
