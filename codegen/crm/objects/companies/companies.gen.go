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

const (
	Oauth2Scopes = "oauth2.Scopes"
)

// Defines values for CreateCompanyJSONBodyAssociationsTypesAssociationCategory.
const (
	HUBSPOTDEFINED    CreateCompanyJSONBodyAssociationsTypesAssociationCategory = "HUBSPOT_DEFINED"
	INTEGRATORDEFINED CreateCompanyJSONBodyAssociationsTypesAssociationCategory = "INTEGRATOR_DEFINED"
	Search            CreateCompanyJSONBodyAssociationsTypesAssociationCategory = "Search"
	USERDEFINED       CreateCompanyJSONBodyAssociationsTypesAssociationCategory = "USER_DEFINED"
)

// GetCompaniesParams defines parameters for GetCompanies.
type GetCompaniesParams struct {
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

// CreateCompanyJSONBody defines parameters for CreateCompany.
type CreateCompanyJSONBody struct {
	// Associations List of associations for the company.
	Associations *[]struct {
		// To Target object details for the association.
		To *struct {
			// Id Target object ID.
			Id *string `json:"id,omitempty"`
		} `json:"to,omitempty"`
		Types *[]struct {
			// AssociationCategory Category of the association.
			AssociationCategory *CreateCompanyJSONBodyAssociationsTypesAssociationCategory `json:"associationCategory,omitempty"`

			// AssociationTypeId ID of the association type.
			AssociationTypeId *int32 `json:"associationTypeId,omitempty"`
		} `json:"types,omitempty"`
	} `json:"associations,omitempty"`

	// ObjectWriteTraceId Trace ID for object write operations.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs of company properties.
	Properties map[string]string `json:"properties"`
}

// CreateCompanyJSONBodyAssociationsTypesAssociationCategory defines parameters for CreateCompany.
type CreateCompanyJSONBodyAssociationsTypesAssociationCategory string

// SearchCompanyJSONBody defines parameters for SearchCompany.
type SearchCompanyJSONBody struct {
	After      *string       `json:"after,omitempty"`
	Limit      *int          `json:"limit,omitempty"`
	Properties *[]string     `json:"properties,omitempty"`
	Query      *string       `json:"query,omitempty"`
	Schema     *FilterGroups `json:"schema,omitempty"`
	Sorts      *[]string     `json:"sorts,omitempty"`
}

// GetCompanyByIdParams defines parameters for GetCompanyById.
type GetCompanyByIdParams struct {
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

// UpdateCompanyJSONBody defines parameters for UpdateCompany.
type UpdateCompanyJSONBody struct {
	// ObjectWriteTraceId Unique trace ID for the operation.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs representing the deal properties to update.
	Properties map[string]string `json:"properties"`
}

// CreateCompanyJSONRequestBody defines body for CreateCompany for application/json ContentType.
type CreateCompanyJSONRequestBody CreateCompanyJSONBody

// SearchCompanyJSONRequestBody defines body for SearchCompany for application/json ContentType.
type SearchCompanyJSONRequestBody SearchCompanyJSONBody

// UpdateCompanyJSONRequestBody defines body for UpdateCompany for application/json ContentType.
type UpdateCompanyJSONRequestBody UpdateCompanyJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of companies
	// (GET /crm/v3/objects/companies)
	GetCompanies(ctx echo.Context, params GetCompaniesParams) error
	// Create a new company
	// (POST /crm/v3/objects/companies)
	CreateCompany(ctx echo.Context) error
	// Search for companies by email
	// (POST /crm/v3/objects/companies/search)
	SearchCompany(ctx echo.Context) error
	// Delete a company
	// (DELETE /crm/v3/objects/companies/{companyId})
	DeleteCompanyById(ctx echo.Context, companyId string) error
	// Get Company Details
	// (GET /crm/v3/objects/companies/{companyId})
	GetCompanyById(ctx echo.Context, companyId int64, params GetCompanyByIdParams) error
	// Update a company
	// (PATCH /crm/v3/objects/companies/{companyId})
	UpdateCompany(ctx echo.Context, companyId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCompanies converts echo context to params.
func (w *ServerInterfaceWrapper) GetCompanies(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"crm.objects.companies.read"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCompaniesParams
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
	err = w.Handler.GetCompanies(ctx, params)
	return err
}

// CreateCompany converts echo context to params.
func (w *ServerInterfaceWrapper) CreateCompany(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"crm.objects.companies.read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateCompany(ctx)
	return err
}

// SearchCompany converts echo context to params.
func (w *ServerInterfaceWrapper) SearchCompany(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"crm.objects.companies.read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchCompany(ctx)
	return err
}

// DeleteCompanyById converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCompanyById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "companyId" -------------
	var companyId string

	err = runtime.BindStyledParameterWithOptions("simple", "companyId", ctx.Param("companyId"), &companyId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter companyId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"crm.objects.companies.read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteCompanyById(ctx, companyId)
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

	ctx.Set(Oauth2Scopes, []string{"crm.objects.companies.read"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCompanyByIdParams
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
	err = w.Handler.GetCompanyById(ctx, companyId, params)
	return err
}

// UpdateCompany converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCompany(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "companyId" -------------
	var companyId string

	err = runtime.BindStyledParameterWithOptions("simple", "companyId", ctx.Param("companyId"), &companyId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter companyId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"crm.objects.companies.read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateCompany(ctx, companyId)
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

	router.GET(baseURL+"/crm/v3/objects/companies", wrapper.GetCompanies)
	router.POST(baseURL+"/crm/v3/objects/companies", wrapper.CreateCompany)
	router.POST(baseURL+"/crm/v3/objects/companies/search", wrapper.SearchCompany)
	router.DELETE(baseURL+"/crm/v3/objects/companies/:companyId", wrapper.DeleteCompanyById)
	router.GET(baseURL+"/crm/v3/objects/companies/:companyId", wrapper.GetCompanyById)
	router.PATCH(baseURL+"/crm/v3/objects/companies/:companyId", wrapper.UpdateCompany)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"o4M4v9P/hDwd4sX14n8BAAD//109ULBqNAAA",
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
