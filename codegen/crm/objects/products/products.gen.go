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

// Defines values for CreateProductJSONBodyAssociationsAssociationCategory.
const (
	HUBSPOTDEFINED    CreateProductJSONBodyAssociationsAssociationCategory = "HUBSPOT_DEFINED"
	INTEGRATORDEFINED CreateProductJSONBodyAssociationsAssociationCategory = "INTEGRATOR_DEFINED"
	Search            CreateProductJSONBodyAssociationsAssociationCategory = "Search"
	USERDEFINED       CreateProductJSONBodyAssociationsAssociationCategory = "USER_DEFINED"
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
		// AssociationCategory Category of the association.
		AssociationCategory *CreateProductJSONBodyAssociationsAssociationCategory `json:"associationCategory,omitempty"`

		// AssociationTypeId ID of the association type.
		AssociationTypeId *int32 `json:"associationTypeId,omitempty"`

		// To Target object details for the association.
		To *struct {
			// Id Target object ID.
			Id *string `json:"id,omitempty"`
		} `json:"to,omitempty"`
	} `json:"associations,omitempty"`

	// ObjectWriteTraceId Trace ID for object write operations.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs of product properties.
	Properties map[string]string `json:"properties"`
}

// CreateProductJSONBodyAssociationsAssociationCategory defines parameters for CreateProduct.
type CreateProductJSONBodyAssociationsAssociationCategory string

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
	Properties         struct {
		// HsSku The new sku of the product.
		HsSku *string `json:"hs_sku,omitempty"`

		// Name The new name of the product.
		Name *string `json:"name,omitempty"`

		// Price The new price of the product.
		Price *int `json:"price,omitempty"`
	} `json:"properties"`
}

// CreateProductJSONRequestBody defines body for CreateProduct for application/json ContentType.
type CreateProductJSONRequestBody CreateProductJSONBody

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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateProduct(ctx)
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
	router.DELETE(baseURL+"/crm/v3/objects/products/:productId", wrapper.DeleteProductById)
	router.GET(baseURL+"/crm/v3/objects/products/:productId", wrapper.GetProductById)
	router.PATCH(baseURL+"/crm/v3/objects/products/:productId", wrapper.UpdateProduct)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"9X8CAAD//14PFiDkLAAA",
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
