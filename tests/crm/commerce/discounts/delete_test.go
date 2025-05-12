package discounts_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

func TestDeleteDiscount(t *testing.T) {
	crmClient := testsutil.GetClient()

	discountId := "410308216595"

	ct := crmClient.Discounts()

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
