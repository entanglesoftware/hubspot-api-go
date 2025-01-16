package deals_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

// TestCreateSchema tests the creation of a schema in HubSpot CRM
func TestCreateSchema(t *testing.T) {
	// Fetch the access token from the environment
	token := os.Getenv("HS_ACCESS_TOKEN")

	if token == "" {
		t.Skip("HS_ACCESS_TOKEN is not set. Skipping test.")
	}

	// Initialize the configuration
	config := configuration.Configuration{
		AccessToken:            token,
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}

	hsClient := hubspot.NewClient(config)

	// Define the schema payload

	name := "email"
	label := "Primary Field"
	description := "The primary field of the custom object."
	proptype := "string"
	fieldType := "text"
	pluralLabel := "My objects"
	singularLabel := "My object"

	schema := schemas.CreateCustomObjectSchemaJSONRequestBody{
		RequiredProperties: []string{
			"email",
		},
		Name: "new_object",
		AssociatedObjects: []string{
			"contact",
		},
		PrimaryDisplayProperty: "email",
		Properties: schemas.SchemaProperties{
			{
				Name:        &name,
				Label:       &label,
				Description: &description,
				Type:        &proptype,
				FieldType:   &fieldType,
			},
		},
		Labels: schemas.SchemaLabels{
			Plural:   &pluralLabel,
			Singular: &singularLabel,
		},
	}

	// Serialize the product properties to JSON
	body, err := json.Marshal(schema)
	if err != nil {
		log.Fatalf("Error serializing product properties: %v", err)
	}

	contentType := "application/json"

	// Make the API call to create the schema
	response, err := hsClient.Crm().SchemaItems().CreateCustomObjectSchemaWithBodyWithResponse(context.Background(), contentType, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		t.Logf("Schema created successfully")
	} else {
		t.Fatalf("Test failed with status code %d: %v", response.StatusCode(), response)
	}
}
