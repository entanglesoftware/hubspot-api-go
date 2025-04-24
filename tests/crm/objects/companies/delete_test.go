package companies_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

// TestDeleteCompanyById fetches a page of companies
func TestDeleteCompanyById(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	ct := hsClient.Crm().Companies()

	response, err := ct.DeleteCompanyByIdWithResponse(context.Background(), "28189124426")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Company Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
