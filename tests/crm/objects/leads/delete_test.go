package leads_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

// TestDeleteLeadById fetches a page of leads
func TestDeleteLeadById(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Leads()

	response, err := ct.DeleteLeadByIdWithResponse(context.Background(), "396774476665")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Lead Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
