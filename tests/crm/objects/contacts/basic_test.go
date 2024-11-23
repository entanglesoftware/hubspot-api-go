package contacts_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
	"testing"
)

// Test setup constants
const (
	BaseURL = "https://api.hubapi.com"
)

// TestGetContacts fetches a page of contacts
func TestGetContacts(t *testing.T) {
	// Fetch the access token from the environment
	// token := os.Getenv("HUBSPOT_ACCESS_TOKEN")
	// if token == "" {
	// 	t.Skip("HUBSPOT_ACCESS_TOKEN is not set. Skipping test.")
	// }

	token := "CN_a0PK1MhINAAEAQAAAASIAAACAARjNmfAWIP3d9iIo3ouMAjIUqxypWYO8RkRYiXzWONETlc5hTa86MAAAAEEAAAAAAAAAAAAAAAAAgAAAAAAAAAAAACAAAAAOAOARAAAAAABAAOABAABQAkIU7ynvGBuHweL1FMav0q-tW3Sz4NhKA25hMVIAWgBgAA"

	hubspotClient := hubspot.Client{}

	// Initialize the client
	hubspotClient.SetAccessToken(token)

	// Make the API call
	response, err := hubspotClient.crm.GetContacts(context.Background(), 10, "", nil, nil, nil, false)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}
}
