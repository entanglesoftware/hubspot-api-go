package companies_test

import (
	"context"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

// TestDeleteCompanyById fetches a page of companies
func TestDeleteCompanyById(t *testing.T) {
	// Fetch the access token from the environment
	token := os.Getenv("HS_ACCESS_TOKEN")

	if token == "" {
		t.Skip("HS_ACCESS_TOKEN is not set. Skipping test.")
	}

	// Correctly initialize the struct with the proper syntax
	config := configuration.Configuration{
		AccessToken:            token,
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}

	hsClient := hubspot.NewClient(config)

	// Initialize the client
	hsClient.SetAccessToken(token)

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
