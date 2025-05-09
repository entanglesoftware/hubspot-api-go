package associations_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

func TestDeleteAssociationsDetails(t *testing.T) {
	crmClient := testsutil.GetClient()
	objectType := "contact"
	objectId := "100260047027"
	toObjectType := "companies"
	toObjectId := "28292686395"

	ct := crmClient.Details()

	response, err := ct.DeleteAssociationWithResponse(context.Background(), objectType, objectId, toObjectType, toObjectId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Associations Details Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
