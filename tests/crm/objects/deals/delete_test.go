package deals_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

// TestDeleteDealById fetches a page of deals
func TestDeleteDealById(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	ct := hsClient.Crm().Deals()

	response, err := ct.DeleteDealByIdWithResponse(context.Background(), "31738621965")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Deal Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
