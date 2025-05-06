package carts_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/carts"
	"github.com/entanglesoftware/hubspot-api-go/configuration"

	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

func TestSearchCarts(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	// Make the API call

	propertyName := "hs_cart_name"
	value := "Cart Name"
	limit := 10

	ct := crm.Carts()

	body := carts.SearchCartsJSONRequestBody{
		Limit: &limit,
		FilterGroups: []carts.FilterGroups{{
			Filters: []carts.Filter{{
				Operator:     carts.FilterOperator("EQ"),
				PropertyName: propertyName,
				Value:        value,
			}},
		}},
	}

	response, err := ct.SearchCartsWithResponse(context.Background(), body)
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
