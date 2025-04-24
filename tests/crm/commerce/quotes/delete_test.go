package quotes_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestDeleteQuote(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	quoteId := "14923556151"

	ct := hsClient.Crm().Quotes()

	response, err := ct.DeleteQuoteByIdWithResponse(context.Background(), quoteId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Quote Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
