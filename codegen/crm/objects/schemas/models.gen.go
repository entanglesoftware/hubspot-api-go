// Package schemas provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package schemas

// SchemaLabels defines model for SchemaLabels.
type SchemaLabels struct {
	// Plural Plural label of the custom object.
	Plural *string `json:"plural,omitempty"`

	// Singular Singular label of the custom object.
	Singular *string `json:"singular,omitempty"`
}

// SchemaProperties defines model for SchemaProperties.
type SchemaProperties = []struct {
	// Description Description of the property.
	Description *string `json:"description,omitempty"`

	// DisplayOrder Display order of the property.
	DisplayOrder *int `json:"displayOrder,omitempty"`

	// FieldType Field type of the property.
	FieldType *string `json:"fieldType,omitempty"`

	// FormField Indicates if the property is a form field.
	FormField *bool `json:"formField,omitempty"`

	// GroupName Group name of the property.
	GroupName *string `json:"groupName,omitempty"`

	// HasUniqueValue Indicates if the property must have unique values.
	HasUniqueValue *bool `json:"hasUniqueValue,omitempty"`

	// Hidden Whether the property is hidden.
	Hidden *bool `json:"hidden,omitempty"`

	// IsPrimaryDisplayLabel Indicates if this property is the primary display label.
	IsPrimaryDisplayLabel *bool `json:"isPrimaryDisplayLabel,omitempty"`

	// Label Label of the property.
	Label *string `json:"label,omitempty"`

	// Name Name of the property.
	Name *string `json:"name,omitempty"`

	// NumberDisplayHint Display hint for number fields.
	NumberDisplayHint *string `json:"numberDisplayHint,omitempty"`

	// OptionSortStrategy Controls how the property options will be sorted in the HubSpot UI.
	OptionSortStrategy *string      `json:"optionSortStrategy,omitempty"`
	Options            *interface{} `json:"options,omitempty"`

	// ReferencedObjectType Referenced object type for the property.
	ReferencedObjectType *string `json:"referencedObjectType,omitempty"`

	// SearchableInGlobalSearch Indicates if the property is searchable in global search.
	SearchableInGlobalSearch *bool `json:"searchableInGlobalSearch,omitempty"`

	// ShowCurrencySymbol Whether to show a currency symbol.
	ShowCurrencySymbol *bool `json:"showCurrencySymbol,omitempty"`

	// TextDisplayHint Display hint for text fields.
	TextDisplayHint *string `json:"textDisplayHint,omitempty"`

	// Type Type of the property.
	Type *string `json:"type,omitempty"`
}

// SchemaPropertiesOptions defines model for SchemaPropertiesOptions.
type SchemaPropertiesOptions = []struct {
	// Description Description of the option.
	Description *string `json:"description,omitempty"`

	// DisplayOrder Display order of the option.
	DisplayOrder *int `json:"displayOrder,omitempty"`

	// Hidden Whether the option is hidden.
	Hidden *bool `json:"hidden,omitempty"`

	// Label Label of the option.
	Label *string `json:"label,omitempty"`

	// Value Value of the option.
	Value *string `json:"value,omitempty"`
}

// SchemaRequestBody defines model for SchemaRequestBody.
type SchemaRequestBody struct {
	// AssociatedObjects Objects that can be associated with the custom object.
	AssociatedObjects []string `json:"associatedObjects"`

	// Description Description of the custom object.
	Description *string     `json:"description,omitempty"`
	Labels      interface{} `json:"labels"`

	// Name Name of the custom object.
	Name string `json:"name"`

	// PrimaryDisplayProperty Primary display property of the custom object.
	PrimaryDisplayProperty string           `json:"primaryDisplayProperty"`
	Properties             SchemaProperties `json:"properties"`

	// RequiredProperties Properties that are required for the custom object.
	RequiredProperties []string `json:"requiredProperties"`

	// SearchableProperties Properties that are searchable for the custom object.
	SearchableProperties *[]string `json:"searchableProperties,omitempty"`

	// SecondaryDisplayProperties Secondary display properties for the custom object.
	SecondaryDisplayProperties *[]string `json:"secondaryDisplayProperties,omitempty"`
}
