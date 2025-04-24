package contacts_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

// TestDeleteContactById fetches a page of contacts
func TestDeleteContactById(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	ct := hsClient.Crm().Contacts()

	response, err := ct.DeleteContactByIdWithResponse(context.Background(), "87484938935")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Contact Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
