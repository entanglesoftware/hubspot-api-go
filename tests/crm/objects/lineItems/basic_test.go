package lineItems_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/lineItems"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestGetLineItems(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	limit := 10

	// Make the API call
	LineItemParams := lineItems.GetLineItemsParams{
		Limit: &limit,
	}

	ct := hsClient.Crm().LineItems()

	response, err := ct.GetLineItemsWithResponse(context.Background(), &LineItemParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
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

func TestGetLineItemById(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	// Make the API call
	LineItemByIdParam := lineItems.GetLineItemByIdParams{}

	ct := hsClient.Crm().LineItems()

	response, err := ct.GetLineItemByIdWithResponse(context.Background(), "27907650925", &LineItemByIdParam)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			t.Fatalf("Response contains no results")
		}

		for key, result := range response.JSON200.Properties {
			t.Logf("Key: %s, Value: %+v\n", key, result)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestSaveLineItem(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	ct := hsClient.Crm().LineItems()

	body := lineItems.CreateLineItemJSONRequestBody{
		Properties: map[string]string{
			"hs_product_id": "18080770487",
			"quantity":      "10",
			"amount":        "2000",
		},
	}

	response, err := ct.CreateLineItemWithResponse(context.Background(), body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		if response.JSON201 == nil || response.JSON201.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON201.Properties != nil {
			t.Logf("Properties: %s", response.JSON201.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
