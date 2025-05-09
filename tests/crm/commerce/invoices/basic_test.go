package invoices_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/invoices"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

// TestGetInvoices fetches a page of invoices
func TestGetInvoices(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	limit := 10

	// Make the API call
	invoicesParams := invoices.GetInvoicesParams{
		Limit: &limit,
	}

	ct := crm.Invoices()

	response, err := ct.GetInvoicesWithResponse(context.Background(), &invoicesParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range response.JSON200.Results {
			t.Logf("%+v\n", result)
			t.Log("-----")

			// Assuming Properties is a map of key-value pairs
			if result.Properties != nil {
				for key, value := range result.Properties {
					t.Logf("Key: %s, Value: %+v\n", key, value)
				}
			} else {
				t.Log("No properties found.")
			}
			t.Log("-----")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

// TestGetInvoiceById fetches a page of invoices
func TestGetInvoiceById(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	// Make the API call
	invoiceByIdParam := invoices.GetInvoiceByIdParams{}

	ct := crm.Invoices()

	response, err := ct.GetInvoiceByIdWithResponse(context.Background(), 408547726686, &invoiceByIdParam)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			t.Fatalf("Response contains no results")
		}

		for key, result := range response.JSON200.Properties {
			t.Logf("Key: %s, Value: %+v\n", key, result)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestSaveInvoices(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	// Initialize a variable of type Invoices
	invoice := invoices.CreateInvoiceJSONRequestBody{
		Properties: map[string]string{
			"hs_title":    "Invoices 2",
			"hs_language": "en",
			"hs_currency": "USD",
		},
	}

	ct := crm.Invoices()
	response, err := ct.CreateInvoiceWithResponse(context.Background(), invoice)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		if response.JSON201 == nil || response.JSON201.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON201.Properties != nil {
			t.Logf("Properties: %s", response.JSON201.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
