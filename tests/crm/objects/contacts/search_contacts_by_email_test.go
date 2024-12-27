package contacts_test

import (
	"context"
	"encoding/json"
	"os"
	"test/codegen/crm/objects/contacts"
	"test/configuration"
	"test/hubspot"
	"testing"
)

func TestSearchContactsByEmail(t *testing.T) {
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

	// Make the API call

	jsonInput := `{
		"filters": [
			{
				"propertyName": "email",
				"operator": "EQ",
				"value": "johndoe57hy@example.com"
			}
		],
		"limit": 1
	}`
	contactByEmailParam := contacts.SearchContactsByEmailParams{}

	var body contacts.SearchContactsByEmailJSONRequestBody

	if err := json.Unmarshal([]byte(jsonInput), &body); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
		return
	}

	ct := hsClient.Crm().Contacts().Contacts

	response, err := ct.SearchContactsByEmailWithResponse(context.Background(), &contactByEmailParam, body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	var result struct {
		Total int `json:"total"`
	}

	if err := json.Unmarshal(response.Body, &result); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if result.Total == 0 {
		t.Fatalf("Response contains no results")
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range *response.JSON200.Results {
			t.Logf("%+v\n", result)
			t.Log("-----")

			// Assuming Properties is a map of key-value pairs
			if result.Properties != nil {
				for key, value := range *result.Properties {
					t.Logf("Key: %s, Value: %+v\n", key, value)
				}
			} else {
				t.Log("No properties found.")
			}
			t.Log("-----")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
