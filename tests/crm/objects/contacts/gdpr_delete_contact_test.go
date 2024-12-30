package contacts_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

func TestGDPRDeleteContact(t *testing.T) {
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

	// Define the payload
	payload := map[string]string{
		"idProperty": "email",
		"objectId":   "example2@example.com",
	}

	// Convert the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error serializing contact properties: %v", err)
	}

	ct := hsClient.Crm().Contacts()

	contentType := "application/json"

	response, err := ct.GdprDeleteContactWithBodyWithResponse(context.Background(), contentType, bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatalf("API call failed: %v", response)
	}
	// t.Logf("API call failed: %v", response.JSON204)
	if response.StatusCode() == 204 {
		t.Logf("Contact Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
