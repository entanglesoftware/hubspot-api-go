package schemas_test

import (
	"context"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

func TestGetSchemas(t *testing.T) {
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

	// Make the API call
	ticketParams := schemas.GetObjectSchemasParams{}

	ct := hsClient.Crm().SchemaItems()

	response, err := ct.GetObjectSchemasWithResponse(context.Background(), &ticketParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range *response.JSON200.Results {
			t.Log("-----\n")
			t.Logf("%+v\n", result)
			t.Log("-----\n")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestGetExistingObjectSchema(t *testing.T) {
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

	// Make the API call
	objectType := "contacts"

	ct := hsClient.Crm().SchemaItems()

	response, err := ct.GetExistingObjectSchemaWithResponse(context.Background(), objectType)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range *response.JSON200 {
			t.Log("-----\n")
			t.Logf("%+v\n", result)
			t.Log("-----\n")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

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
			"company",
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

	// Make the API call to create the schema
	response, err := hsClient.Crm().SchemaItems().CreateCustomObjectSchemaWithResponse(context.Background(), schema)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		t.Logf("Schema created successfully")
	} else {
		t.Fatalf("Test failed with status code %d: %v", response.StatusCode(), response)
	}
}
