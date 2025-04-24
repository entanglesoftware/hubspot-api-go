package deals_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/deals"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestSearchDeals(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	// Make the API call

	propertyName := "dealname"
	value := "New Deal 1"
	limit := 10

	ct := hsClient.Crm().Deals()

	body := deals.SearchDealsJSONRequestBody{
		Limit: &limit,
		FilterGroups: []deals.FilterGroups{{
			Filters: []deals.Filter{{
				Operator:     deals.FilterOperator("EQ"),
				PropertyName: propertyName,
				Value:        value,
			}},
		}},
	}

	response, err := ct.SearchDealsWithResponse(context.Background(), body)
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
