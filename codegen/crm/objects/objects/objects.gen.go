// Package objects provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package objects

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

// Defines values for CreateObjectJSONBodyAssociationsTypesAssociationCategory.
const (
	HUBSPOTDEFINED    CreateObjectJSONBodyAssociationsTypesAssociationCategory = "HUBSPOT_DEFINED"
	INTEGRATORDEFINED CreateObjectJSONBodyAssociationsTypesAssociationCategory = "INTEGRATOR_DEFINED"
	Search            CreateObjectJSONBodyAssociationsTypesAssociationCategory = "Search"
	USERDEFINED       CreateObjectJSONBodyAssociationsTypesAssociationCategory = "USER_DEFINED"
)

// GetObjectsParams defines parameters for GetObjects.
type GetObjectsParams struct {
	// Limit Number of objects to retrieve.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Archived Include archived objects.
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty"`
}

// CreateObjectJSONBody defines parameters for CreateObject.
type CreateObjectJSONBody struct {
	// Associations List of associations for the lead.
	Associations *[]struct {
		// To Target object details for the association.
		To *struct {
			// Id Target object ID.
			Id *string `json:"id,omitempty"`
		} `json:"to,omitempty"`
		Types *[]struct {
			// AssociationCategory Category of the association.
			AssociationCategory *CreateObjectJSONBodyAssociationsTypesAssociationCategory `json:"associationCategory,omitempty"`

			// AssociationTypeId ID of the association type.
			AssociationTypeId *int32 `json:"associationTypeId,omitempty"`
		} `json:"types,omitempty"`
	} `json:"associations,omitempty"`

	// Properties Key-value pairs of lead properties.
	Properties *map[string]string `json:"properties,omitempty"`
}

// CreateObjectJSONBodyAssociationsTypesAssociationCategory defines parameters for CreateObject.
type CreateObjectJSONBodyAssociationsTypesAssociationCategory string

// SearchObjectsJSONBody defines parameters for SearchObjects.
type SearchObjectsJSONBody struct {
	After        *string       `json:"after,omitempty"`
	FilterGroups *FilterGroups `json:"filterGroups,omitempty"`
	Limit        *int          `json:"limit,omitempty"`
	Properties   *[]string     `json:"properties,omitempty"`
	Query        *string       `json:"query,omitempty"`
	Sorts        *[]string     `json:"sorts,omitempty"`
}

// GetObjectByTypeAndIdParams defines parameters for GetObjectByTypeAndId.
type GetObjectByTypeAndIdParams struct {
	// PropertiesWithHistory Properties to fetch with history.
	PropertiesWithHistory *[]string `form:"propertiesWithHistory,omitempty" json:"propertiesWithHistory,omitempty"`

	// Associations Associations to include in the response.
	Associations *[]string `form:"associations,omitempty" json:"associations,omitempty"`

	// Archived Whether to include archived objects.
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty"`

	// IdProperty The property to use for ID lookups.
	IdProperty *string `form:"idProperty,omitempty" json:"idProperty,omitempty"`
}

// UpdateObjectJSONBody defines parameters for UpdateObject.
type UpdateObjectJSONBody struct {
	// ObjectWriteTraceId Unique trace ID for the operation.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs representing the deal properties to update.
	Properties map[string]string `json:"properties"`
}

// UpdateObjectParams defines parameters for UpdateObject.
type UpdateObjectParams struct {
	// IdProperty The property to identify the object (e.g., "email").
	IdProperty *string `form:"idProperty,omitempty" json:"idProperty,omitempty"`
}

// CreateObjectJSONRequestBody defines body for CreateObject for application/json ContentType.
type CreateObjectJSONRequestBody CreateObjectJSONBody

// SearchObjectsJSONRequestBody defines body for SearchObjects for application/json ContentType.
type SearchObjectsJSONRequestBody SearchObjectsJSONBody

// UpdateObjectJSONRequestBody defines body for UpdateObject for application/json ContentType.
type UpdateObjectJSONRequestBody UpdateObjectJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get CRM Objects
	// (GET /crm/v3/objects/{objectType})
	GetObjects(ctx echo.Context, objectType string, params GetObjectsParams) error
	// Create an object in HubSpot CRM
	// (POST /crm/v3/objects/{objectType})
	CreateObject(ctx echo.Context, objectType string) error
	// Search HubSpot CRM objects
	// (POST /crm/v3/objects/{objectType}/search)
	SearchObjects(ctx echo.Context, objectType string) error
	// Delete an object by objectType and objectId
	// (DELETE /crm/v3/objects/{objectType}/{objectId})
	DeleteObject(ctx echo.Context, objectType string, objectId string) error
	// Retrieve object details with associations and properties
	// (GET /crm/v3/objects/{objectType}/{objectId})
	GetObjectByTypeAndId(ctx echo.Context, objectType string, objectId string, params GetObjectByTypeAndIdParams) error
	// Update a CRM object
	// (PATCH /crm/v3/objects/{objectType}/{objectId})
	UpdateObject(ctx echo.Context, objectType string, objectId string, params UpdateObjectParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetObjects converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjects(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "objectType" -------------
	var objectType string

	err = runtime.BindStyledParameterWithOptions("simple", "objectType", ctx.Param("objectType"), &objectType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectType: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"objects"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "archived" -------------

	err = runtime.BindQueryParameter("form", true, false, "archived", ctx.QueryParams(), &params.Archived)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter archived: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetObjects(ctx, objectType, params)
	return err
}

// CreateObject converts echo context to params.
func (w *ServerInterfaceWrapper) CreateObject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "objectType" -------------
	var objectType string

	err = runtime.BindStyledParameterWithOptions("simple", "objectType", ctx.Param("objectType"), &objectType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectType: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"objects"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateObject(ctx, objectType)
	return err
}

// SearchObjects converts echo context to params.
func (w *ServerInterfaceWrapper) SearchObjects(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "objectType" -------------
	var objectType string

	err = runtime.BindStyledParameterWithOptions("simple", "objectType", ctx.Param("objectType"), &objectType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectType: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"objects"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchObjects(ctx, objectType)
	return err
}

// DeleteObject converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteObject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "objectType" -------------
	var objectType string

	err = runtime.BindStyledParameterWithOptions("simple", "objectType", ctx.Param("objectType"), &objectType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectType: %s", err))
	}

	// ------------- Path parameter "objectId" -------------
	var objectId string

	err = runtime.BindStyledParameterWithOptions("simple", "objectId", ctx.Param("objectId"), &objectId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"objects"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteObject(ctx, objectType, objectId)
	return err
}

// GetObjectByTypeAndId converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectByTypeAndId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "objectType" -------------
	var objectType string

	err = runtime.BindStyledParameterWithOptions("simple", "objectType", ctx.Param("objectType"), &objectType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectType: %s", err))
	}

	// ------------- Path parameter "objectId" -------------
	var objectId string

	err = runtime.BindStyledParameterWithOptions("simple", "objectId", ctx.Param("objectId"), &objectId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"objects"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectByTypeAndIdParams
	// ------------- Optional query parameter "propertiesWithHistory" -------------

	err = runtime.BindQueryParameter("form", true, false, "propertiesWithHistory", ctx.QueryParams(), &params.PropertiesWithHistory)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter propertiesWithHistory: %s", err))
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
	err = w.Handler.GetObjectByTypeAndId(ctx, objectType, objectId, params)
	return err
}

// UpdateObject converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateObject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "objectType" -------------
	var objectType string

	err = runtime.BindStyledParameterWithOptions("simple", "objectType", ctx.Param("objectType"), &objectType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectType: %s", err))
	}

	// ------------- Path parameter "objectId" -------------
	var objectId string

	err = runtime.BindStyledParameterWithOptions("simple", "objectId", ctx.Param("objectId"), &objectId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter objectId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"objects"})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdateObjectParams
	// ------------- Optional query parameter "idProperty" -------------

	err = runtime.BindQueryParameter("form", true, false, "idProperty", ctx.QueryParams(), &params.IdProperty)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter idProperty: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateObject(ctx, objectType, objectId, params)
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

	router.GET(baseURL+"/crm/v3/objects/:objectType", wrapper.GetObjects)
	router.POST(baseURL+"/crm/v3/objects/:objectType", wrapper.CreateObject)
	router.POST(baseURL+"/crm/v3/objects/:objectType/search", wrapper.SearchObjects)
	router.DELETE(baseURL+"/crm/v3/objects/:objectType/:objectId", wrapper.DeleteObject)
	router.GET(baseURL+"/crm/v3/objects/:objectType/:objectId", wrapper.GetObjectByTypeAndId)
	router.PATCH(baseURL+"/crm/v3/objects/:objectType/:objectId", wrapper.UpdateObject)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"dv2/AAAA//+5gW4AgisAAA==",
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
