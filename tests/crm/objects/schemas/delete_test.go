package schemas_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
)

// TestDeleteSchema tests the creation of a schema in HubSpot CRM
func TestDeleteSchema(t *testing.T) {
	crmClient := testsutil.GetClient()
	archived := false
	deleteSchemaParams := schemas.DeleteSchemaParams{
		Archived: &archived,
	}

	objectType := "2-39502275"

	// Make the API call to create the schema
	response, err := crmClient.SchemaItems().DeleteSchemaWithResponse(context.Background(), objectType, &deleteSchemaParams)
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
	crmClient := testsutil.GetClient()

	objectType := "2-39502275"
	associationIdentifier := "608"

	// Make the API call to create the schema
	response, err := crmClient.SchemaItems().DeleteAssociationWithResponse(context.Background(), objectType, associationIdentifier)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Association Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
