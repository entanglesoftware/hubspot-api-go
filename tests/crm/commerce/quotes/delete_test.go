package quotes_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

func TestDeleteQuote(t *testing.T) {
	crmClient := testsutil.GetClient()

	quoteId := "14923556151"

	ct := crmClient.Quotes()

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
