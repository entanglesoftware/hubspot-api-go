// Package deals provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package deals

import (
	"time"
)

// Defines values for FilterOperator.
const (
	CONTAINSTOKEN    FilterOperator = "CONTAINS_TOKEN"
	EQ               FilterOperator = "EQ"
	GT               FilterOperator = "GT"
	GTE              FilterOperator = "GTE"
	HASPROPERTY      FilterOperator = "HAS_PROPERTY"
	LT               FilterOperator = "LT"
	LTE              FilterOperator = "LTE"
	NEQ              FilterOperator = "NEQ"
	NOTCONTAINSTOKEN FilterOperator = "NOT_CONTAINS_TOKEN"
	NOTHASPROPERTY   FilterOperator = "NOT_HAS_PROPERTY"
)

// AssociationResponse defines model for AssociationResponse.
type AssociationResponse struct {
	// Id The ID of the associated object.
	Id *string `json:"id,omitempty"`

	// Type The type of association.
	Type *string `json:"type,omitempty"`
}

// DealResponse defines model for DealResponse.
type DealResponse struct {
	// Archived Indicates if the deal is archived.
	Archived bool `json:"archived,omitempty"`

	// ArchivedAt When the deal was archived.
	ArchivedAt time.Time `json:"archivedAt,omitempty"`

	// Associations A map of associated objects.
	Associations map[string]ObjectAssociationsResponse `json:"associations,omitempty"`

	// CreatedAt When the deal was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// Id Unique identifier for the deal.
	Id string `json:"id,omitempty"`

	// Properties A key-value map of the deal's properties.
	Properties map[string]string `json:"properties,omitempty"`

	// PropertiesWithHistory A map of the deal's properties including historical values.
	PropertiesWithHistory map[string][]PropertyHistory `json:"propertiesWithHistory,omitempty"`

	// UpdatedAt When the deal was last updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// DealsResponse defines model for DealsResponse.
type DealsResponse struct {
	Paging  *Paging        `json:"paging,omitempty"`
	Results []DealResponse `json:"results,omitempty"`
}

// Filter defines model for Filter.
type Filter struct {
	// HighValue A high-value filter criterion.
	HighValue *string `json:"highValue,omitempty"`

	// Operator defines model for Filter.Operator
	Operator *FilterOperator `json:"operator,omitempty"`

	// PropertyName The property name to filter by.
	PropertyName *string `json:"propertyName,omitempty"`

	// Value A single value to match for the property.
	Value *string `json:"value,omitempty"`

	// Values List of values to match for the property.
	Values *[]string `json:"values,omitempty"`
}

// FilterOperator defines model for Filter.Operator
type FilterOperator string

// FilterGroups defines model for FilterGroups.
type FilterGroups = []struct {
	Filters *[]Filter `json:"Filters,omitempty"`
}

// ObjectAssociationsResponse defines model for ObjectAssociationsResponse.
type ObjectAssociationsResponse struct {
	Paging  *Paging                `json:"paging,omitempty"`
	Results *[]AssociationResponse `json:"results,omitempty"`
}

// Paging defines model for Paging.
type Paging struct {
	Next PagingNext `json:"next,omitempty"`
}

// PagingNext defines model for PagingNext.
type PagingNext struct {
	// After The cursor token for the next page of results.
	After string `json:"after,omitempty"`

	// Link The link for the next page of results.
	Link string `json:"link,omitempty"`
}

// PropertyHistory defines model for PropertyHistory.
type PropertyHistory struct {
	// SourceId The source ID of the historical property value.
	SourceId string `json:"sourceId,omitempty"`

	// SourceLabel The source label for the historical property.
	SourceLabel string `json:"sourceLabel,omitempty"`

	// SourceType The source type of the historical property value.
	SourceType string `json:"sourceType,omitempty"`

	// Timestamp When the property value was set.
	Timestamp time.Time `json:"timestamp,omitempty"`

	// UpdatedByUserId The user ID who updated the property.
	UpdatedByUserId int `json:"updatedByUserId,omitempty"`

	// Value The historical value of the property.
	Value string `json:"value,omitempty"`
}

// Archived defines model for Archived.
type Archived = bool

// Associations defines model for Associations.
type Associations = []string

// Properties defines model for Properties.
type Properties = []string

// PropertiesWithHistory defines model for PropertiesWithHistory.
type PropertiesWithHistory = []string
