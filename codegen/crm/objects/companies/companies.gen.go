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
	After        *string        `json:"after,omitempty"`
	FilterGroups []FilterGroups `json:"filterGroups,omitempty"`
	Limit        *int           `json:"limit,omitempty"`
	Properties   *[]string      `json:"properties,omitempty"`
	Query        *string        `json:"query,omitempty"`
	Sorts        *[]string      `json:"sorts,omitempty"`
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
	"HJKMDuL8Vv8T8nSIF1eL/wUAAP//eWCUKZo1AAA=",
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
