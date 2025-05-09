package invoices_test

import (
	"context"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/invoices"
)

func TestUpdateInvoice(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Invoices()

	invoiceId := "408555114028"

	body := invoices.UpdateInvoiceJSONRequestBody{
		Properties: map[string]string{
			"hs_title": "Invoice Update",
		},
	}

	response, err := ct.UpdateInvoiceWithResponse(context.Background(), invoiceId, body)
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
