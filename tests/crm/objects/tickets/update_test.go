package tickets_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/tickets"
)

func TestUpdateTicket(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Initialize a variable of type Ticket
	subject := "Update Ticket"
	hs_pipeline_stage := "2"
	hs_pipeline := "0"

	ticketId := "18791135765"

	ct := crmClient.Tickets()

	body := tickets.UpdateTicketJSONRequestBody{
		Properties: map[string]string{
			"subject":           subject,
			"hs_pipeline_stage": hs_pipeline_stage,
			"hs_pipeline":       hs_pipeline,
		},
	}

	response, err := ct.UpdateTicketWithResponse(context.Background(), ticketId, body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON200.Properties != nil {
			t.Logf("Properties: %s", response.JSON200.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
