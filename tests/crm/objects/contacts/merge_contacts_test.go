package contacts_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"test/configuration"
	"test/hubspot"
	"testing"
)

func TestMergeContacts(t *testing.T) {
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
		"objectIdToMerge": "87484939431",
		"primaryObjectId": "83910492845",
	}

	// Convert the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error serializing contact properties: %v", err)
	}

	ct := hsClient.Crm().Contacts().Contacts

	contentType := "application/json"

	response, err := ct.MergeContactsWithBodyWithResponse(context.Background(), contentType, bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatalf("API call failed: %v", response)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			t.Fatalf("Response contains no results")
		}

		for key, result := range *response.JSON200.Properties {
			t.Logf("Key: %s, Value: %+v\n", key, result)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
