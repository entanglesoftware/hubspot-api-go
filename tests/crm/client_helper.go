package crm

import (
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

// GetTestHubSpotClient initializes and returns the HubSpot client for tests
func GetTestHubSpotClient(t *testing.T) *hubspot.Client {
	t.Helper()

	token := os.Getenv("HS_ACCESS_TOKEN")
	if token == "" {
		t.Skip("HS_ACCESS_TOKEN is not set. Skipping test.")
	}

	config := configuration.Configuration{
		AccessToken:            token,
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}

	client := hubspot.NewClient(config)
	client.SetAccessToken(token)

	return client
}
