package users_test

import (
	"context"
	"encoding/json"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/users"
)

func TestSearchUsers(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Make the API call

	propertyName := "hs_family_name"
	value := "parmar"
	limit := 10

	ct := crmClient.Users()

	body := users.SearchUsersJSONRequestBody{
		Limit: &limit,
		FilterGroups: []users.FilterGroups{{
			Filters: []users.Filter{{
				Operator:     users.FilterOperator("EQ"),
				PropertyName: propertyName,
				Value:        value,
			}},
		}},
		Properties: &[]string{"hs_job_title", "hs_availability_status", "hs_working_hours"},
	}

	response, err := ct.SearchUsersWithResponse(context.Background(), body)
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
