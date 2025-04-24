package lineItems_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

// TestDeleteLineItemsById fetches a page of lineItems
func TestDeleteLineItemsById(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	ct := hsClient.Crm().LineItems()

	response, err := ct.DeleteLineItemByIdWithResponse(context.Background(), "27907650925")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("LineItems Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
