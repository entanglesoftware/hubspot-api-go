package objects_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestDeleteObject(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	objectType := "contacts"

	objectId := "87481797267"

	ct := hsClient.Crm().Objects()

	response, err := ct.DeleteObjectWithResponse(context.Background(), objectType, objectId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Objects Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
