package invoices_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestDeleteInvoice(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	invoiceId := "408555114028"

	ct := hsClient.Crm().Invoices()

	response, err := ct.DeleteInvoiceByIdWithResponse(context.Background(), invoiceId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Invoice Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
