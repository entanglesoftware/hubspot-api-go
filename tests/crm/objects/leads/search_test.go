package leads_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/leads"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestSearchLeads(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	// Make the API call

	propertyName := "hs_pipeline_stage"
	value := "new-stage-id"
	limit := 10

	ct := hsClient.Crm().Leads()

	body := leads.SearchLeadsJSONRequestBody{
		Limit: &limit,
		FilterGroups: []leads.FilterGroups{{
			Filters: []leads.Filter{{
				Operator:     leads.FilterOperator("EQ"),
				PropertyName: propertyName,
				Value:        value,
			}},
		}},
	}

	response, err := ct.SearchLeadsWithResponse(context.Background(), body)
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
