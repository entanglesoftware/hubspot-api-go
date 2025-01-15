package leads_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/leads"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

func TestSearchLeads(t *testing.T) {
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

	leadByEmailParam := leads.SearchLeadsParams{}

	propertyName := "hs_pipeline_stage"
	operator := "EQ"
	value := "new-stage-id"
	limit := 10

	filters := []struct {
		Operator     *string `json:"operator"`
		PropertyName *string `json:"propertyName"`
		Value        *string `json:"value"`
	}{
		{
			Operator:     &operator,
			PropertyName: &propertyName,
			Value:        &value,
		},
	}

	filterGroups := []struct {
		Filters *[]struct {
			Operator     *string `json:"operator"`
			PropertyName *string `json:"propertyName"`
			Value        *string `json:"value"`
		} `json:"filters"`
	}{
		{
			Filters: &filters,
		},
	}

	body := struct {
		Limit        *int `json:"limit"`
		FilterGroups *[]struct {
			Filters *[]struct {
				Operator     *string `json:"operator"`
				PropertyName *string `json:"propertyName"`
				Value        *string `json:"value"`
			} `json:"filters"`
		} `json:"filterGroups"`
	}{
		Limit:        &limit,
		FilterGroups: &filterGroups,
	}

	// Convert body to JSON
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Error marshalling body: %v", err)
	}

	// Convert JSON to io.Reader
	bodyReader := bytes.NewReader(bodyJSON)

	contentType := "application/json"

	ct := hsClient.Crm().Leads()

	response, err := ct.SearchLeadsWithBodyWithResponse(context.Background(), &leadByEmailParam, contentType, bodyReader)
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
	t.Logf("Total result : %+v\n", result.Total)

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range response.JSON200.Results {
			t.Logf("%+v\n", result)
			t.Log("-----")

			// Assuming Properties is a map of key-value pairs
			if result.Properties != nil {
				for key, value := range result.Properties {
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
