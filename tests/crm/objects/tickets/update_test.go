package tickets_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/tickets"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

func TestUpdateTicket(t *testing.T) {
	// Fetch the access token from the environment
	token := os.Getenv("HS_ACCESS_TOKEN")

	if token == "" {
		t.Skip("HS_ACCESS_TOKEN is not set. Skipping test.")
	}

	// Correctly initialize the struct with the proper syntax
	config := configuration.Configuration{
		AccessToken:            token,
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}

	hsClient := hubspot.NewClient(config)

	// Initialize the client
	hsClient.SetAccessToken(token)

	// Initialize a variable of type Ticket
	subject := "Update Ticket"
	hs_pipeline_stage := "2"
	hs_pipeline := "0"

	ticket := tickets.UpdateTicketJSONBody{
		Properties: map[string]string{
			"subject":           subject,
			"hs_pipeline_stage": hs_pipeline_stage,
			"hs_pipeline":       hs_pipeline,
		},
	}

	ticketId := "18816298665"

	// Serialize the ticket properties to JSON
	body, err := json.Marshal(ticket)
	if err != nil {
		log.Fatalf("Error serializing ticket properties: %v", err)
	}

	contentType := "application/json"

	ct := hsClient.Crm().Tickets()

	response, err := ct.UpdateTicketWithBodyWithResponse(context.Background(), ticketId, contentType, bytes.NewReader(body))
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
