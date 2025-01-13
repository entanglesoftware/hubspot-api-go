// Package deals provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package deals

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

// Defines values for CreateDealJSONBodyAssociationsTypesAssociationCategory.
const (
	HUBSPOTDEFINED    CreateDealJSONBodyAssociationsTypesAssociationCategory = "HUBSPOT_DEFINED"
	INTEGRATORDEFINED CreateDealJSONBodyAssociationsTypesAssociationCategory = "INTEGRATOR_DEFINED"
	Search            CreateDealJSONBodyAssociationsTypesAssociationCategory = "Search"
	USERDEFINED       CreateDealJSONBodyAssociationsTypesAssociationCategory = "USER_DEFINED"
)

// GetDealsParams defines parameters for GetDeals.
type GetDealsParams struct {
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

// CreateDealJSONBody defines parameters for CreateDeal.
type CreateDealJSONBody struct {
	// Associations List of associations for the deal.
	Associations *[]struct {
		// To Target object details for the association.
		To *struct {
			// Id Target object ID.
			Id *string `json:"id,omitempty"`
		} `json:"to,omitempty"`
		Types *[]struct {
			// AssociationCategory Category of the association.
			AssociationCategory *CreateDealJSONBodyAssociationsTypesAssociationCategory `json:"associationCategory,omitempty"`

			// AssociationTypeId ID of the association type.
			AssociationTypeId *int32 `json:"associationTypeId,omitempty"`
		} `json:"types,omitempty"`
	} `json:"associations,omitempty"`

	// ObjectWriteTraceId Trace ID for object write operations.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs of deal properties.
	Properties map[string]string `json:"properties"`
}

// CreateDealJSONBodyAssociationsTypesAssociationCategory defines parameters for CreateDeal.
type CreateDealJSONBodyAssociationsTypesAssociationCategory string

// SearchDealsJSONBody defines parameters for SearchDeals.
type SearchDealsJSONBody struct {
	After      *string   `json:"after,omitempty"`
	Limit      *int      `json:"limit,omitempty"`
	Properties *[]string `json:"properties,omitempty"`
	Query      *string   `json:"query,omitempty"`
	Schema     *Filters  `json:"schema,omitempty"`
	Sorts      *[]string `json:"sorts,omitempty"`
}

// SearchDealsParams defines parameters for SearchDeals.
type SearchDealsParams struct {
	// Hapikey HubSpot API key
	Hapikey string `form:"hapikey" json:"hapikey"`
}

// GetDealByIdParams defines parameters for GetDealById.
type GetDealByIdParams struct {
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

// UpdateDealJSONBody defines parameters for UpdateDeal.
type UpdateDealJSONBody struct {
	// ObjectWriteTraceId Unique trace ID for the operation.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs representing the deal properties to update.
	Properties map[string]string `json:"properties"`
}

// CreateDealJSONRequestBody defines body for CreateDeal for application/json ContentType.
type CreateDealJSONRequestBody CreateDealJSONBody

// SearchDealsJSONRequestBody defines body for SearchDeals for application/json ContentType.
type SearchDealsJSONRequestBody SearchDealsJSONBody

// UpdateDealJSONRequestBody defines body for UpdateDeal for application/json ContentType.
type UpdateDealJSONRequestBody UpdateDealJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of deals
	// (GET /crm/v3/objects/deals)
	GetDeals(ctx echo.Context, params GetDealsParams) error
	// Create a new deal
	// (POST /crm/v3/objects/deals)
	CreateDeal(ctx echo.Context) error
	// Search for deals by email
	// (POST /crm/v3/objects/deals/search)
	SearchDeals(ctx echo.Context, params SearchDealsParams) error
	// Delete a deal
	// (DELETE /crm/v3/objects/deals/{dealId})
	DeleteDealById(ctx echo.Context, dealId string) error
	// Get Deal Details
	// (GET /crm/v3/objects/deals/{dealId})
	GetDealById(ctx echo.Context, dealId string, params GetDealByIdParams) error
	// Update a deal
	// (PATCH /crm/v3/objects/deals/{dealId})
	UpdateDeal(ctx echo.Context, dealId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDeals converts echo context to params.
func (w *ServerInterfaceWrapper) GetDeals(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"crm.objects.deals.read"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDealsParams
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
	err = w.Handler.GetDeals(ctx, params)
	return err
}

// CreateDeal converts echo context to params.
func (w *ServerInterfaceWrapper) CreateDeal(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"crm.objects.deals.read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateDeal(ctx)
	return err
}

// SearchDeals converts echo context to params.
func (w *ServerInterfaceWrapper) SearchDeals(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"crm.objects.deals.read"})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchDealsParams
	// ------------- Required query parameter "hapikey" -------------

	err = runtime.BindQueryParameter("form", true, true, "hapikey", ctx.QueryParams(), &params.Hapikey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hapikey: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchDeals(ctx, params)
	return err
}

// DeleteDealById converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteDealById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "dealId" -------------
	var dealId string

	err = runtime.BindStyledParameterWithOptions("simple", "dealId", ctx.Param("dealId"), &dealId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter dealId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"crm.objects.deals.read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteDealById(ctx, dealId)
	return err
}

// GetDealById converts echo context to params.
func (w *ServerInterfaceWrapper) GetDealById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "dealId" -------------
	var dealId string

	err = runtime.BindStyledParameterWithOptions("simple", "dealId", ctx.Param("dealId"), &dealId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter dealId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"crm.objects.deals.read"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDealByIdParams
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
	err = w.Handler.GetDealById(ctx, dealId, params)
	return err
}

// UpdateDeal converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateDeal(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "dealId" -------------
	var dealId string

	err = runtime.BindStyledParameterWithOptions("simple", "dealId", ctx.Param("dealId"), &dealId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter dealId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"crm.objects.deals.read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateDeal(ctx, dealId)
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

	router.GET(baseURL+"/crm/v3/objects/deals", wrapper.GetDeals)
	router.POST(baseURL+"/crm/v3/objects/deals", wrapper.CreateDeal)
	router.POST(baseURL+"/crm/v3/objects/deals/search", wrapper.SearchDeals)
	router.DELETE(baseURL+"/crm/v3/objects/deals/:dealId", wrapper.DeleteDealById)
	router.GET(baseURL+"/crm/v3/objects/deals/:dealId", wrapper.GetDealById)
	router.PATCH(baseURL+"/crm/v3/objects/deals/:dealId", wrapper.UpdateDeal)

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
