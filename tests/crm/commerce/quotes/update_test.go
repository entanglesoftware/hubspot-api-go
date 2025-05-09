package quotes_test

import (
	"context"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/quotes"
)

func TestUpdateQuote(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Quotes()

	quoteId := "14923556151"

	body := quotes.UpdateQuoteJSONRequestBody{
		Properties: map[string]string{
			"hs_title": "Deal Update",
		},
	}

	response, err := ct.UpdateQuoteWithResponse(context.Background(), quoteId, body)
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
