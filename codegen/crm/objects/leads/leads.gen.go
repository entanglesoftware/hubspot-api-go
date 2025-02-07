// Package leads provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package leads

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

// Defines values for CreateLeadJSONBodyAssociationsTypesAssociationCategory.
const (
	HUBSPOTDEFINED    CreateLeadJSONBodyAssociationsTypesAssociationCategory = "HUBSPOT_DEFINED"
	INTEGRATORDEFINED CreateLeadJSONBodyAssociationsTypesAssociationCategory = "INTEGRATOR_DEFINED"
	Search            CreateLeadJSONBodyAssociationsTypesAssociationCategory = "Search"
	USERDEFINED       CreateLeadJSONBodyAssociationsTypesAssociationCategory = "USER_DEFINED"
)

// GetLeadsParams defines parameters for GetLeads.
type GetLeadsParams struct {
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

// CreateLeadJSONBody defines parameters for CreateLead.
type CreateLeadJSONBody struct {
	// Associations List of associations for the lead.
	Associations []struct {
		// To Target object details for the association.
		To *struct {
			// Id Target object ID.
			Id *string `json:"id,omitempty"`
		} `json:"to,omitempty"`
		Types *[]struct {
			// AssociationCategory Category of the association.
			AssociationCategory *CreateLeadJSONBodyAssociationsTypesAssociationCategory `json:"associationCategory,omitempty"`

			// AssociationTypeId ID of the association type.
			AssociationTypeId *int32 `json:"associationTypeId,omitempty"`
		} `json:"types,omitempty"`
	} `json:"associations"`

	// ObjectWriteTraceId Trace ID for object write operations.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs of lead properties.
	Properties map[string]string `json:"properties"`
}

// CreateLeadJSONBodyAssociationsTypesAssociationCategory defines parameters for CreateLead.
type CreateLeadJSONBodyAssociationsTypesAssociationCategory string

// SearchLeadsJSONBody defines parameters for SearchLeads.
type SearchLeadsJSONBody struct {
	After        *string       `json:"after,omitempty"`
	FilterGroups *FilterGroups `json:"filterGroups,omitempty"`
	Limit        *int          `json:"limit,omitempty"`
	Properties   *[]string     `json:"properties,omitempty"`
	Query        *string       `json:"query,omitempty"`
	Sorts        *[]string     `json:"sorts,omitempty"`
}

// GetLeadByIdParams defines parameters for GetLeadById.
type GetLeadByIdParams struct {
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

// UpdateLeadJSONBody defines parameters for UpdateLead.
type UpdateLeadJSONBody struct {
	// ObjectWriteTraceId Unique trace ID for the operation.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs representing the lead properties to update.
	Properties map[string]string `json:"properties"`
}

// CreateLeadJSONRequestBody defines body for CreateLead for application/json ContentType.
type CreateLeadJSONRequestBody CreateLeadJSONBody

// SearchLeadsJSONRequestBody defines body for SearchLeads for application/json ContentType.
type SearchLeadsJSONRequestBody SearchLeadsJSONBody

// UpdateLeadJSONRequestBody defines body for UpdateLead for application/json ContentType.
type UpdateLeadJSONRequestBody UpdateLeadJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of leads
	// (GET /crm/v3/objects/leads)
	GetLeads(ctx echo.Context, params GetLeadsParams) error
	// Create a new lead
	// (POST /crm/v3/objects/leads)
	CreateLead(ctx echo.Context) error
	// Search for leads by email
	// (POST /crm/v3/objects/leads/search)
	SearchLeads(ctx echo.Context) error
	// Delete a lead
	// (DELETE /crm/v3/objects/leads/{leadId})
	DeleteLeadById(ctx echo.Context, leadId string) error
	// Get Lead Details
	// (GET /crm/v3/objects/leads/{leadId})
	GetLeadById(ctx echo.Context, leadId string, params GetLeadByIdParams) error
	// Update a lead
	// (PATCH /crm/v3/objects/leads/{leadId})
	UpdateLead(ctx echo.Context, leadId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetLeads converts echo context to params.
func (w *ServerInterfaceWrapper) GetLeads(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"leads"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLeadsParams
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
	err = w.Handler.GetLeads(ctx, params)
	return err
}

// CreateLead converts echo context to params.
func (w *ServerInterfaceWrapper) CreateLead(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"leads"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateLead(ctx)
	return err
}

// SearchLeads converts echo context to params.
func (w *ServerInterfaceWrapper) SearchLeads(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"leads"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchLeads(ctx)
	return err
}

// DeleteLeadById converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteLeadById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "leadId" -------------
	var leadId string

	err = runtime.BindStyledParameterWithOptions("simple", "leadId", ctx.Param("leadId"), &leadId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter leadId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"leads"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteLeadById(ctx, leadId)
	return err
}

// GetLeadById converts echo context to params.
func (w *ServerInterfaceWrapper) GetLeadById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "leadId" -------------
	var leadId string

	err = runtime.BindStyledParameterWithOptions("simple", "leadId", ctx.Param("leadId"), &leadId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter leadId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"leads"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLeadByIdParams
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
	err = w.Handler.GetLeadById(ctx, leadId, params)
	return err
}

// UpdateLead converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateLead(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "leadId" -------------
	var leadId string

	err = runtime.BindStyledParameterWithOptions("simple", "leadId", ctx.Param("leadId"), &leadId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter leadId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"leads"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateLead(ctx, leadId)
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

	router.GET(baseURL+"/crm/v3/objects/leads", wrapper.GetLeads)
	router.POST(baseURL+"/crm/v3/objects/leads", wrapper.CreateLead)
	router.POST(baseURL+"/crm/v3/objects/leads/search", wrapper.SearchLeads)
	router.DELETE(baseURL+"/crm/v3/objects/leads/:leadId", wrapper.DeleteLeadById)
	router.GET(baseURL+"/crm/v3/objects/leads/:leadId", wrapper.GetLeadById)
	router.PATCH(baseURL+"/crm/v3/objects/leads/:leadId", wrapper.UpdateLead)

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
