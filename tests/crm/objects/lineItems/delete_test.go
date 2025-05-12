package lineItems_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

// TestDeleteLineItemsById fetches a page of lineItems
func TestDeleteLineItemsById(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.LineItems()

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
