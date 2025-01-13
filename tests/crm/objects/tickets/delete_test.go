package tickets_test

import (
	"context"
	"os"
	"testing"

	"test/configuration"
	"test/hubspot"
)

// TestDeleteTicketById fetches a page of tickets
func TestDeleteTicketById(t *testing.T) {
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

	ct := hsClient.Crm().Tickets()

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
