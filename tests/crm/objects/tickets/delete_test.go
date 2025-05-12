package tickets_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

// TestDeleteTicketById fetches a page of tickets
func TestDeleteTicketById(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Tickets()

	response, err := ct.DeleteTicketByIdWithResponse(context.Background(), "18816298665")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Ticket Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
