package taxes_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/taxes"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestUpdateTax(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	ct := hsClient.Crm().Taxes()

	taxId := "404043653204"

	body := taxes.UpdateTaxJSONRequestBody{
		Properties: map[string]string{
			"hs_label":      "Taxes 123",
			"hs_value":      "300",
			"hs_type":       "FIXED",
			"hs_sort_order": "1",
		},
	}

	response, err := ct.UpdateTaxWithResponse(context.Background(), taxId, body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			fmt.Println("No data found.")
			return
		}

		properties := response.JSON200.Properties
		if properties == nil {
			fmt.Println("No properties found.")
			return
		}
		for key, value := range properties {
			fmt.Printf("%s: %s\n", key, value)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
