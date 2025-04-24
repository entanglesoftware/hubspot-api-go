package taxes_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestDeleteTax(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	taxId := "416498264080"

	ct := hsClient.Crm().Taxes()

	response, err := ct.DeleteTaxByIdWithResponse(context.Background(), taxId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Tax Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
