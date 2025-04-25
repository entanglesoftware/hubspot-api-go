package schemas_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

// TestCreateAssociation tests the creation of a schema in HubSpot CRM
func TestCreateAssociation(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

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
	response, err := crm.SchemaItems().CreateAssociationWithResponse(context.Background(), objectType, association)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		t.Logf("Association created successfully")
	} else {
		t.Fatalf("Test failed with status code %d: %v", response.StatusCode(), response)
	}
}
