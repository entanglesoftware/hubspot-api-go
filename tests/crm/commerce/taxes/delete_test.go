package taxes_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

func TestDeleteTax(t *testing.T) {
	crmClient := testsutil.GetClient()

	taxId := "416498264080"

	ct := crmClient.Taxes()

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
