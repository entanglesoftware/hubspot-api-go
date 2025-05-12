package associations_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

func TestDeleteAssociationsSchemas(t *testing.T) {
	crmClient := testsutil.GetClient()

	fromObjectType := "contact"
	toObjectType := "companies"
	associationTypeId := 77

	ct := crmClient.Schemas()

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
