package discounts_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestDeleteDiscount(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	discountId := "410308216595"

	ct := hsClient.Crm().Discounts()

	response, err := ct.DeleteDiscountByIdWithResponse(context.Background(), discountId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Discount Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
