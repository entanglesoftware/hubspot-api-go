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

// GetProductsParams defines parameters for GetProducts.
type GetProductsParams struct {
	// Limit Maximum number of results per page.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// After Cursor token for the next page of results.
	After *string `form:"after,omitempty" json:"after,omitempty"`

	// Properties Comma-separated list of properties to include in the response.
	// If a specified property is not present, it will be ignored.
	Properties *[]string `form:"properties,omitempty" json:"properties,omitempty"`

	// PropertiesWithHistory Comma-separated list of properties to include with their historical values.
	// Historical data reduces the maximum number of objects returned per request.
	PropertiesWithHistory *[]string `form:"propertiesWithHistory,omitempty" json:"propertiesWithHistory,omitempty"`

	// Associations Comma-separated list of object types to retrieve associated IDs for.
	// Nonexistent associations will be ignored.
	Associations *[]string `form:"associations,omitempty" json:"associations,omitempty"`

	// Archived Include only archived results.
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of products
	// (GET /crm/v3/objects/products)
	GetProducts(ctx echo.Context, params GetProductsParams) error
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

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xYTW/jNhP+KwTfF+jFX9ugFwF7cLttY6BJje0u9rBZLGhpZDGRSIYcOTYC/feClGRR",
	"H3bsOil6MQyKnHnmmU/ymYYyU1KAQEODZ2rCBDLm/jYfvs+NkSFnyKX4CEZJYcDuUFoq0MjB7eeR/Y3A",
	"hJoru5UG9FMCZPGByJhgAoRVYiAicnUPIU7oiOJOAQ2oQc3Fmhb1wpAk+8XKYg2cAQl7EbRUYmV6tvzp",
	"Fj2LzGGTFFtbkcEz/b+GmAb0f9NG0rTiylv6viwPFCOqweRpySlHyMwZQobYboxiWrPdi1Yu98jbFgnY",
	"4tn23NpDfY0juh2v5dgujs0DV2PpnMXSsZJcIGgaoM5hCNhtBaMNjsXu0JDrw1wbqQnKBxAktv8SINYY",
	"otjaBUVFeD8gToeZcvEwrN5+eSO1r0SrllEe4uFQZjpM+AYGcnQhIh4yBEN4maahFMhCJNyQ+pRn3krK",
	"FJg4A2YtZI595V8SEC2lT6ytNZY6Y0gDGjGEMfIMLnCwVzhKUqKIl1uXLbJOzI4jpaQYdQydk4wpv3jt",
	"i6AfO+eHgAYr6mRqq/2vzuxQ8f8s+GMOhEcgkMcc9D6DKkSX5Gonug84stdcuk55gN14w9Icavd48H4w",
	"pNFyiZMaKV84JtfcoNS7Y7DP7RjV6V0tutctDkbjIXMJF2GaR1ysSeKE8pClxDF1ERW5is6K15QZJNWh",
	"Vw7alzpoWVH/c9NBt9J3ff165rdCqme9kbkOYXFg6Cu/erOfF0WVnF0ZTpcUgVLLH2wF6VEYqd2xrz4D",
	"UC4H8eng1FphqIfXNyPDJoNBlqkjqdXW5jLMAL56N6gS9ufdZwP6UIjkBrQNkKdE1hnewuhxYaWvQZ+B",
	"wNk3rLdbz2qnXB4L/ZSyS1zEso9kvly4gIwBw8SWWVXVGxJrmZHrfPWXkrayIsfUiqxWSF2XyHy5oCO6",
	"AW1KiZsrm8JSgWCK04BeTWaTGR1RxTBxGTsNdTbdXE2rsWNaa7Tf1jBQkX+z2MA00J44JsSVPTfuECai",
	"Tou0/90363X6O2AN1wHRLAMEbWjwtavrhm15lmdE5NkKtDdZEwXaTdtWPGxZpiwdP80ssTSgjzloW/QE",
	"yyxLKc+4Za0snKVJMctTpMG7WS+eimLUxfHLOReNPRx6cz9/dzN//54OwypvNj6s3qW1B0RmGRsbsKzZ",
	"1Ei5Qave69EoqzYNhJf5rauuMLkTi5gwYhSEdvSKmsznhgiJRGkwIHBEOJInnqZkBYSvhdQQTe5EY5vx",
	"OGwPKSbPMmY7A/213OpBo/sM/Epjrg06HkbUtvPqL2SMp/SbzRDYqlRGQIOYpQaGGWzJbmjcd88DzwjN",
	"AGRw5zxlCx29lG6XCJgA1wPj0Z24btYihoxoiPLQSkjsnNkN9CofiQbMtbC+Ak00POZg8ARfdGbKY24p",
	"cSfV5pd89A9842P5V91UcuiarHOUBtQcNq03p8UHYzN6cidupYAtNwgC/cckc3ImzDsXyT7praumx3QE",
	"LDWurIcPgOZ0kjsC34DbRRXcUqS7/R18sNwdg1m/MwzW4Opg5ymhKL65kdiVLmfKj7NZ+QQprI/cVUmp",
	"lIfO/Om9sXifPQXnzc7eNb3oXY2WZXvzYqu6j5iJewNrXP1xH2N+tSi7XVEUxd8BAAD///Ka6t1cFQAA",
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