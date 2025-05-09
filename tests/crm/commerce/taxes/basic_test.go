package taxes_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/taxes"
)

// TestGetTaxes fetches a page of taxes
func TestGetTaxes(t *testing.T) {
	crmClient := testsutil.GetClient()

	limit := 10

	// Make the API call
	taxesParams := taxes.GetTaxesParams{
		Limit: &limit,
	}

	ct := crmClient.Taxes()
	response, err := ct.GetTaxesWithResponse(context.Background(), &taxesParams)
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

// TestGetTaxById fetches a page of taxes
func TestGetTaxById(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Make the API call
	invoiceByIdParam := taxes.GetTaxByIdParams{}

	ct := crmClient.Taxes()

	response, err := ct.GetTaxByIdWithResponse(context.Background(), 404043653204, &invoiceByIdParam)
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

func TestSaveTaxes(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Initialize a variable of type Taxes
	invoice := taxes.CreateTaxJSONRequestBody{
		Properties: map[string]string{
			"hs_label":      "Taxes",
			"hs_value":      "20",
			"hs_sort_order": "1",
			"hs_type":       "FIXED",
		},
	}

	ct := crmClient.Taxes()
	response, err := ct.CreateTaxWithResponse(context.Background(), invoice)
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
