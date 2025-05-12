package companies_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

// TestDeleteCompanyById fetches a page of companies
func TestDeleteCompanyById(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Companies()

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
