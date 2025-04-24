package orders_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestDeleteOrder(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	taxId := "417774243908"

	ct := hsClient.Crm().Orders()

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
