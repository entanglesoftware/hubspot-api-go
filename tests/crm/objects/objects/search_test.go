package objects_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestSearchLineItems(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	// Make the API call

	// Make the API call

	propertyName := "lastname"
	value := "Doe1"
	limit := 10

	objectType := "contacts"

	ct := hsClient.Crm().Objects()

	body := objects.SearchObjectsJSONRequestBody{
		FilterGroups: []objects.FilterGroups{{
			Filters: []objects.Filter{{
				Operator:     objects.FilterOperator("EQ"),
				PropertyName: propertyName,
				Value:        value,
			}},
		}},
		Limit: &limit,
	}

	response, err := ct.SearchObjectsWithResponse(context.Background(), objectType, body)

	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	var result struct {
		Total int `json:"total"`
	}

	if err := json.Unmarshal(response.Body, &result); err != nil {
		t.Logf("Failed to parse response body: %v", err)
	}

	if result.Total == 0 {
		t.Fatalf("Response contains no results")
	}

	t.Logf("Total results: %+v\n", result.Total)

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
