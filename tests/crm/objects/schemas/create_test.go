package schemas_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
)

// TestCreateAssociation tests the creation of a schema in HubSpot CRM
func TestCreateAssociation(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Define the Association payload

	fromObjectTypeId := "2-123456"
	name := "my_object_to_contact"
	toObjectTypeId := "contact"

	association := schemas.CreateAssociationJSONRequestBody{
		FromObjectTypeId: fromObjectTypeId,
		Name:             &name,
		ToObjectTypeId:   toObjectTypeId,
	}

	objectType := "contacts"

	// Make the API call to create the association
	response, err := crmClient.SchemaItems().CreateAssociationWithResponse(context.Background(), objectType, association)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		t.Logf("Association created successfully")
	} else {
		t.Fatalf("Test failed with status code %d: %v", response.StatusCode(), response)
	}
}
