package quotes_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/quotes"
	"github.com/entanglesoftware/hubspot-api-go/configuration"

	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

func TestUpdateQuote(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	ct := crm.Quotes()

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
