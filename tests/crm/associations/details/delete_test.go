package associations_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestDeleteAssociationsDetails(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	objectType := "contact"
	objectId := "100260047027"
	toObjectType := "companies"
	toObjectId := "28292686395"

	ct := hsClient.Crm().Details()

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
