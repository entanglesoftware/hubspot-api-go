package orders_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	crmClient := testsutil.GetClient()

	taxId := "417774243908"

	ct := crmClient.Orders()

	response, err := ct.DeleteOrderByIdWithResponse(context.Background(), taxId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Order Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
