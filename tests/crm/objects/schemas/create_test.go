package schemas_test

import (
	"context"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

// TestCreateAssociation tests the creation of a schema in HubSpot CRM
func TestCreateAssociation(t *testing.T) {
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
	response, err := hsClient.Crm().SchemaItems().CreateAssociationWithResponse(context.Background(), objectType, association)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		t.Logf("Association created successfully")
	} else {
		t.Fatalf("Test failed with status code %d: %v", response.StatusCode(), response)
	}
}
