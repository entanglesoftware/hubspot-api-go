package invoices_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

func TestDeleteInvoice(t *testing.T) {
	crmClient := testsutil.GetClient()

	invoiceId := "408555114028"

	ct := crmClient.Invoices()

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
