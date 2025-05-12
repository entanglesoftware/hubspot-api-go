package discounts_test

import (
	"context"
	"fmt"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/discounts"
)

func TestUpdateDiscount(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Discounts()

	discountId := "410308216595"

	body := discounts.UpdateDiscountJSONRequestBody{
		Properties: map[string]string{
			"hs_label":      "Discounts 123",
			"hs_value":      "300",
			"hs_duration":   "ONCE",
			"hs_type":       "FIXED",
			"hs_sort_order": "1",
		},
	}

	response, err := ct.UpdateDiscountWithResponse(context.Background(), discountId, body)
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
