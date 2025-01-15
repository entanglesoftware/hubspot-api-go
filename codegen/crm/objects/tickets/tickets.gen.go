// Package tickets provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package tickets

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

// Defines values for CreateTicketJSONBodyAssociationsTypesAssociationCategory.
const (
	HUBSPOTDEFINED    CreateTicketJSONBodyAssociationsTypesAssociationCategory = "HUBSPOT_DEFINED"
	INTEGRATORDEFINED CreateTicketJSONBodyAssociationsTypesAssociationCategory = "INTEGRATOR_DEFINED"
	Search            CreateTicketJSONBodyAssociationsTypesAssociationCategory = "Search"
	USERDEFINED       CreateTicketJSONBodyAssociationsTypesAssociationCategory = "USER_DEFINED"
)

// GetTicketsParams defines parameters for GetTickets.
type GetTicketsParams struct {
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

// CreateTicketJSONBody defines parameters for CreateTicket.
type CreateTicketJSONBody struct {
	// Associations List of associations for the ticket.
	Associations *[]struct {
		// To Target object details for the association.
		To *struct {
			// Id Target object ID.
			Id *string `json:"id,omitempty"`
		} `json:"to,omitempty"`
		Types *[]struct {
			// AssociationCategory Category of the association.
			AssociationCategory *CreateTicketJSONBodyAssociationsTypesAssociationCategory `json:"associationCategory,omitempty"`

			// AssociationTypeId ID of the association type.
			AssociationTypeId *int32 `json:"associationTypeId,omitempty"`
		} `json:"types,omitempty"`
	} `json:"associations,omitempty"`

	// ObjectWriteTraceId Trace ID for object write operations.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs of ticket properties.
	Properties map[string]string `json:"properties"`
}

// CreateTicketJSONBodyAssociationsTypesAssociationCategory defines parameters for CreateTicket.
type CreateTicketJSONBodyAssociationsTypesAssociationCategory string

// SearchTicketsJSONBody defines parameters for SearchTickets.
type SearchTicketsJSONBody struct {
	After        *string       `json:"after,omitempty"`
	FilterGroups *FilterGroups `json:"filterGroups,omitempty"`
	Limit        *int          `json:"limit,omitempty"`
	Properties   *[]string     `json:"properties,omitempty"`
	Query        *string       `json:"query,omitempty"`
	Sorts        *[]string     `json:"sorts,omitempty"`
}

// SearchTicketsParams defines parameters for SearchTickets.
type SearchTicketsParams struct {
	// Hapikey HubSpot API key
	Hapikey string `form:"hapikey" json:"hapikey"`
}

// GetTicketByIdParams defines parameters for GetTicketById.
type GetTicketByIdParams struct {
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

// UpdateTicketJSONBody defines parameters for UpdateTicket.
type UpdateTicketJSONBody struct {
	// ObjectWriteTraceId Unique trace ID for the operation.
	ObjectWriteTraceId *string `json:"objectWriteTraceId,omitempty"`

	// Properties Key-value pairs representing the ticket properties to update.
	Properties map[string]string `json:"properties"`
}

// CreateTicketJSONRequestBody defines body for CreateTicket for application/json ContentType.
type CreateTicketJSONRequestBody CreateTicketJSONBody

// SearchTicketsJSONRequestBody defines body for SearchTickets for application/json ContentType.
type SearchTicketsJSONRequestBody SearchTicketsJSONBody

// UpdateTicketJSONRequestBody defines body for UpdateTicket for application/json ContentType.
type UpdateTicketJSONRequestBody UpdateTicketJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of tickets
	// (GET /crm/v3/objects/tickets)
	GetTickets(ctx echo.Context, params GetTicketsParams) error
	// Create a new ticket
	// (POST /crm/v3/objects/tickets)
	CreateTicket(ctx echo.Context) error
	// Search for tickets by email
	// (POST /crm/v3/objects/tickets/search)
	SearchTickets(ctx echo.Context, params SearchTicketsParams) error
	// Delete a ticket
	// (DELETE /crm/v3/objects/tickets/{ticketId})
	DeleteTicketById(ctx echo.Context, ticketId string) error
	// Get Ticket Details
	// (GET /crm/v3/objects/tickets/{ticketId})
	GetTicketById(ctx echo.Context, ticketId string, params GetTicketByIdParams) error
	// Update a ticket
	// (PATCH /crm/v3/objects/tickets/{ticketId})
	UpdateTicket(ctx echo.Context, ticketId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTickets converts echo context to params.
func (w *ServerInterfaceWrapper) GetTickets(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"tickets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTicketsParams
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
	err = w.Handler.GetTickets(ctx, params)
	return err
}

// CreateTicket converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTicket(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"tickets"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateTicket(ctx)
	return err
}

// SearchTickets converts echo context to params.
func (w *ServerInterfaceWrapper) SearchTickets(ctx echo.Context) error {
	var err error

	ctx.Set(Oauth2Scopes, []string{"tickets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchTicketsParams
	// ------------- Required query parameter "hapikey" -------------

	err = runtime.BindQueryParameter("form", true, true, "hapikey", ctx.QueryParams(), &params.Hapikey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hapikey: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchTickets(ctx, params)
	return err
}

// DeleteTicketById converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTicketById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ticketId" -------------
	var ticketId string

	err = runtime.BindStyledParameterWithOptions("simple", "ticketId", ctx.Param("ticketId"), &ticketId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ticketId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"tickets"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTicketById(ctx, ticketId)
	return err
}

// GetTicketById converts echo context to params.
func (w *ServerInterfaceWrapper) GetTicketById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ticketId" -------------
	var ticketId string

	err = runtime.BindStyledParameterWithOptions("simple", "ticketId", ctx.Param("ticketId"), &ticketId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ticketId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"tickets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTicketByIdParams
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
	err = w.Handler.GetTicketById(ctx, ticketId, params)
	return err
}

// UpdateTicket converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTicket(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ticketId" -------------
	var ticketId string

	err = runtime.BindStyledParameterWithOptions("simple", "ticketId", ctx.Param("ticketId"), &ticketId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ticketId: %s", err))
	}

	ctx.Set(Oauth2Scopes, []string{"tickets"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateTicket(ctx, ticketId)
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

	router.GET(baseURL+"/crm/v3/objects/tickets", wrapper.GetTickets)
	router.POST(baseURL+"/crm/v3/objects/tickets", wrapper.CreateTicket)
	router.POST(baseURL+"/crm/v3/objects/tickets/search", wrapper.SearchTickets)
	router.DELETE(baseURL+"/crm/v3/objects/tickets/:ticketId", wrapper.DeleteTicketById)
	router.GET(baseURL+"/crm/v3/objects/tickets/:ticketId", wrapper.GetTicketById)
	router.PATCH(baseURL+"/crm/v3/objects/tickets/:ticketId", wrapper.UpdateTicket)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xbXW/bOtL+KwTfF9gbxU7as3th4FykSdoabZ2cxD1F0QYpI40tNhKpklQSo/B/X/BD",
	"35Qj10m72O1NmkjkcDhffOah+h2HPM04A6YknnzHGREkBQXC/FW9uzoUYUxvIdKPI5ChoJminOEJnrIw",
	"ySNAnCUrRNwwJEDmiZIjHGC4J2mWAJ4sSCIhwFTP+paDWOEAM5ICnuBiHg6wDGNIiV1nQfJElRPVKtNj",
	"rzlPgDC8XgcNDaXkISVaK9nV8oinKdmToDeoIEIJlQrxBeLXXyFUSIuWSHEkQAkKt4CIEwcRmh5LtOBi",
	"9JnNOIN7KhUwVQ7Q66E7miToGhBdMi4gGn1m1cZlbS9tJWWepkSs8ASf2MENsTjAtyTJAU8+YUXDG1D6",
	"UfHbpTYA3GcJjx4wblNkZWCqIDVqOMtKJShb4nVpaiIEWem/pVppD+IFFylu2v1M8AyEorCF1bNyjrY5",
	"dRFEGVIx6NDJOJMw+symC0SQzCCkCwpRMW2FqESMK5QJkMBUgKga7IGmul37V6p1rW/sWTpAkaX+K6MZ",
	"JJTBcH80lngyb3ygKn5NpeJ6e7s55o6qWLuGChQbiTQkCTLGkaPP7HX1LCKKIAFRHmoJMaCU3NM0TxHL",
	"02sQVcZJnWm5YNqrIJCAbzlINcBrrW1tcqDVO3aDn96bddUew7HrQkinGFcJfe6SxRTvRmhTT62ex4Cm",
	"x9oL2je1Emedoqt1j44+SfqNllWrLx4J1TbtIq2AfUkTBaKrfkyX8d/WYe21D5F+uWfciRZmPgoFVSD8",
	"GgRYCyaKi66sCBaUgUQpjyDRVR5ZhUanxZQAA8tTHTYnf+EAz8zPt3Pz4wQH+NXc/NC/vj68uDo7Pz07",
	"OZ9/1ENP51etR0ens/nhdHZxNT99czJzY1oPLz0bKCrfzMSczxllbdRhqZPXGeZ65bXIbZ9pJWXLBGxy",
	"azEpUWFsDKNqq/TL9JwCb11xse83Cx2aK4Oi6pXgeSYbKdgMMjusOeL/BSzwBP/fuBI3dmk47sbtIMXa",
	"aV4Tc2oG1WFBf0ZnZKltMVzHMzthHWAHx35ko75is607zkrNmzticK+23s9MT+quGOD7vSXf0w/35A3N",
	"9riJP5LsZZwyU2OUyMGn2Myp0VSOLFxh6iZbmAupg5ffACvDWG8GZWRpamIN/zYjebiaCWU3/uX1myda",
	"9pHM6jK6dlA3bSt5LkKY9hxR9m3tpKqhjrLOmWqyi33tKm/JNSQb1Uj0iNLcHlV2V2Lee8Y6HYqj9smM",
	"oWgKUpE066rxIQbWqNPucLgjEkkwoEHjFaL0cUoU7GlZO6iSZ1pK9GL1XoLoC5FcgtABchdz5Cb0HVBa",
	"+hLEFhr0HI7zpvGtFZxTdo+FB0ro3MDU/sOBbGjQIxoSBRJRq6tFvLqJKibVtC666+HWKoQcqg2x49bU",
	"MVNf9FEDh7SaaxJF1A5tNn0DT5sNR/M66OCmlGR1LFxi6not3r6kCtCihlrWDX90w/paifeMfssB0QiY",
	"0t25KCukVWiXYtSK7B43dhBi2yU3sHJNgnNOpd0/ZK1P3MVDWV9j2qf1tvCrfZR64GRPKPbs1nX1lC09",
	"zfwOlnBVeGisJkSqonI/csAOqqT/cTi7VeDbfv7xzetHlC24p9s7m5qsXYAKYx0QjgtBC8FT9Dq/vsi4",
	"MhSJMrSEe4KcBdHh2RQH+BaEtPJun7tmm5GM4gl+Ptof7eMAZ0TFxhTjUKTj2+djVxrHBZU5+Y6X4Amc",
	"l1ox3TK6BQ2hY9xjCjIiLGrlse30KWcaOOBXoOYlb1rntj+1V3rX4amcKw09peF1g8r+576fDEpoSpWf",
	"xD7Y7yCS9TroEHPbdBalOvjd18ODd4d//ol7GGDTytTV6pA0/jitbDb2U74/OrFeLbeT0aDQt5xaYKT1",
	"pUlVk2km+p7t71uOjSlgJhBJliUaNlHOxl+l9s33mvm2yukablh3yvWZjeYaB+uCfWQJwJLfPC9vJtoj",
	"DUHEpSd/jgwgQAQxuCthny3FRS4fnb/r5o2dZ7XH2lSGnX3Bo9VWVirjs1NhC5Z1grODkoN1wWqXPWhR",
	"sxPjo3UXHWy89imop8ZlTReo9NBDylMz50QsQRX3RhEoQpNKZIsKHcDINsRNjwcSqOa+agOxVdPjiChY",
	"+i8B3Js2H+yUL2jPC9CIHQf49fsXF2en86vjk5fT2ckxDvD7i5Pz2p/T2fzk1fnh/LR66CMza+voxtfX",
	"4nVJal3utaQGVqBMPX+GfXX1IRJuCE1n33wQVMFckB66Qr/QjaiOAOfFOz0DlRklvWTpo2DcNyXCzQgV",
	"sqoJG9Htem1TmgrdLn6qq3LpRQ+tLdsVitBX3PUdjTIzwvU1HERsldyDrYrJ0Hb3QwwqBpuOYS4VT0HU",
	"e13EBWJcPVnLOy+YFHT3M5vfDa3iZo1+QdNYZbdb/NF7xtYZWzU/D637uwF8igZwcwT+1Fawt54VISHz",
	"MAQpF3mSrNoYzIOntGpkqVsKPK9/mNHT6oylPU11QfNCtr9yENRdntcwmkUYrg26JlLXMWvIL5ASmnyp",
	"GMjP7NxcrRdtU1mpZR7GiEj0ZUGFNGjrS4C+aOu733U75cTZm/gmKrRAYGBDVeiue8wbWPU0JjHJqH3b",
	"PCwCH4rEH0/fn18dnk2v3px89OCkyx9Hqj03P53jd9G6V9zqstBNM1c7uk2s7+2gC2I68LEoHOWkT9h9",
	"wxDLK3mT42ZkXwYPf3Bg3eHbqeSixWNsfSPbxBl+DPCr2i6X9E0Y3053G/DN1FshkyDbpv13+8s0WtuU",
	"T0B5bhiOzXOJSJG6dSIG5dJwNPa24ZZG5fnl+oZmtlpZVrkXq2n0UMKaW5X2+ezkK46syqMijzOi4iqN",
	"i81tzGNvujaC4Y+uRWYcHdnoGKF589ioF2qnXtT2oDVCaU8dtV6iqWyumyUzqJ2VjRZSV8rqgEcpKBIR",
	"RTYwUENd4LN98VHiLtav90z/+sPXZW78qkRxlEvdjRnVbLz5CjqNCrzym276RXTTzmXvFRQULzq2I8xh",
	"RJTFLS1Ib3CbDYsCaPBFVcEo20w12fkl1TS0RNXunTrpYqHko5eqx0AWQygF1ySpOrOgN1jabQc+obnS",
	"374b7KYFtyQgBLhvcouTqkNHeMU/FTFRLfS4OOQ3F/FfyEWUn5H8ZC6id93fXMT/HBdxUce09e+aiuBo",
	"ntT28KzBW/MawlxQtTIHKCe5ip/V/wPJpT7NJIjb4ozNRaI7caUyORmPSUZHcX6t/wl5Osbry/W/AwAA",
	"///J2wneJTQAAA==",
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
