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
	After        *string        `json:"after,omitempty"`
	FilterGroups []FilterGroups `json:"filterGroups,omitempty"`
	Limit        *int           `json:"limit,omitempty"`
	Properties   *[]string      `json:"properties,omitempty"`
	Query        *string        `json:"query,omitempty"`
	Sorts        *[]string      `json:"sorts,omitempty"`
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

// BatchUpsertProductsJSONRequestBody defines body for BatchUpsertProducts for application/json ContentType.
type BatchUpsertProductsJSONRequestBody = BatchProductsUpsertRequest

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
	// Batch create or update products
	// (POST /crm/v3/objects/products/batch/upsert)
	BatchUpsertProducts(ctx echo.Context) error
	// Search for products by email
	// (POST /crm/v3/objects/products/search)
	SearchProducts(ctx echo.Context) error
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

// BatchUpsertProducts converts echo context to params.
func (w *ServerInterfaceWrapper) BatchUpsertProducts(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.BatchUpsertProducts(ctx)
	return err
}

// SearchProducts converts echo context to params.
func (w *ServerInterfaceWrapper) SearchProducts(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"e-commerce"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchProducts(ctx)
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
	router.POST(baseURL+"/crm/v3/objects/products/batch/upsert", wrapper.BatchUpsertProducts)
	router.POST(baseURL+"/crm/v3/objects/products/search", wrapper.SearchProducts)
	router.DELETE(baseURL+"/crm/v3/objects/products/:productId", wrapper.DeleteProductById)
	router.GET(baseURL+"/crm/v3/objects/products/:productId", wrapper.GetProductById)
	router.PATCH(baseURL+"/crm/v3/objects/products/:productId", wrapper.UpdateProduct)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"s4X65bN4iDd3m/8EAAD//9Pa1C+tQQAA",
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
