package associations_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

func TestDeleteAssociationsSchemas(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	fromObjectType := "contact"
	toObjectType := "companies"
	associationTypeId := 77

	ct := crm.Schemas()

	response, err := ct.DeleteAssociationsSchemaWithResponse(context.Background(), fromObjectType, toObjectType, associationTypeId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Associations Schemas Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
