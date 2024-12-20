// Package companies provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package companies

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

// GetCompanyByIdParams defines parameters for GetCompanyById.
type GetCompanyByIdParams struct {
	// Properties A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
	Properties *string `form:"properties,omitempty" json:"properties,omitempty"`

	// Associations A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
	Associations *string `form:"associations,omitempty" json:"associations,omitempty"`

	// Archived Whether to return only results that have been archived. Default value - false.
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty"`

	// IdProperty The name of a property whose values are unique for this object type.
	IdProperty *string `form:"idProperty,omitempty" json:"idProperty,omitempty"`
}

// SearchCompanyJSONRequestBody defines body for SearchCompany for application/json ContentType.
type SearchCompanyJSONRequestBody = SearchParams

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /crm/v3/objects/companies/search)
	SearchCompany(ctx echo.Context) error

	// (GET /crm/v3/objects/companies/{companyId})
	GetCompanyById(ctx echo.Context, companyId int64, params GetCompanyByIdParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// SearchCompany converts echo context to params.
func (w *ServerInterfaceWrapper) SearchCompany(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchCompany(ctx)
	return err
}

// GetCompanyById converts echo context to params.
func (w *ServerInterfaceWrapper) GetCompanyById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "companyId" -------------
	var companyId int64

	err = runtime.BindStyledParameterWithOptions("simple", "companyId", ctx.Param("companyId"), &companyId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter companyId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCompanyByIdParams
	// ------------- Optional query parameter "properties" -------------

	err = runtime.BindQueryParameter("form", true, false, "properties", ctx.QueryParams(), &params.Properties)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter properties: %s", err))
	}

	// ------------- Optional query parameter "associations" -------------

	err = runtime.BindQueryParameter("form", true, false, "associations", ctx.QueryParams(), &params.Associations)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter associations: %s", err))
	}

	// ------------- Optional query parameter "archived" -------------

	err = runtime.BindQueryParameter("form", true, false, "archived", ctx.QueryParams(), &params.Archived)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter archived: %s", err))
	}

	// ------------- Optional query parameter "idProperty" -------------

	err = runtime.BindQueryParameter("form", true, false, "idProperty", ctx.QueryParams(), &params.IdProperty)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter idProperty: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCompanyById(ctx, companyId, params)
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

	router.POST(baseURL+"/crm/v3/objects/companies/search", wrapper.SearchCompany)
	router.GET(baseURL+"/crm/v3/objects/companies/:companyId", wrapper.GetCompanyById)

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
