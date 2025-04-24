package lineItems_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/lineItems"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestUpdateLineItem(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	// Initialize a variable of type LineItem
	quantity := "5"
	amount := "220"

	ticketId := "27907650998"

	ct := hsClient.Crm().LineItems()

	body := lineItems.UpdateLineItemJSONRequestBody{
		Properties: map[string]string{
			"quantity": quantity,
			"amount":   amount,
		},
	}

	response, err := ct.UpdateLineItemWithResponse(context.Background(), ticketId, body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON200.Properties != nil {
			t.Logf("Properties: %s", response.JSON200.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
