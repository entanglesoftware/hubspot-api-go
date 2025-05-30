package contacts_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

func TestSearchContactByEmail(t *testing.T) {
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

	propertyName := "email"
	email := "johndoe12@example.com"
	limit := 10

	body := contacts.SearchContactsJSONRequestBody{
		Limit: &limit,
		FilterGroups: []contacts.FilterGroups{{
			Filters: []contacts.Filter{{
				Operator:     contacts.FilterOperator("EQ"),
				PropertyName: propertyName,
				Value:        email,
			}},
		}},
	}

	ct := hsClient.Crm().Contacts()

	response, err := ct.SearchContactsWithResponse(context.Background(), body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() != 200 {
		t.Fatalf("API call failed with code: %v", response.StatusCode())
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

	if result.Total > 0 && response.JSON200 != nil && response.JSON200.Results != nil {
		fmt.Printf("Total Result : %+v\n", result.Total)
		conatct := response.JSON200.Results[0]
		fmt.Printf("%+v\n", conatct)
	} else {
		fmt.Printf("Email %s does not exist in HubSpot.\n", email)
	}
}
