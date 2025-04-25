package schemas_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

// TestDeleteSchema tests the creation of a schema in HubSpot CRM
func TestDeleteSchema(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)
	archived := false
	deleteSchemaParams := schemas.DeleteSchemaParams{
		Archived: &archived,
	}

	objectType := "2-39502275"

	// Make the API call to create the schema
	response, err := crm.SchemaItems().DeleteSchemaWithResponse(context.Background(), objectType, &deleteSchemaParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Schema Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

// TestAssociationSchema tests the creation of a schema in HubSpot CRM
func TestAssociationSchema(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	objectType := "2-39502275"
	associationIdentifier := "608"

	// Make the API call to create the schema
	response, err := crm.SchemaItems().DeleteAssociationWithResponse(context.Background(), objectType, associationIdentifier)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Association Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
