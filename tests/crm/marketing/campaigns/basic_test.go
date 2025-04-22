package campaigns_test

import (
	"context"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/marketing/campaigns"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

// TestGetCampaigns fetches a page of discounts
func TestGetCampaigns(t *testing.T) {
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

	campaignGuid := "a4226822-7b2d-4cc1-90ca-7f0d1844be6d"

	// Make the API call
	discountsParams := campaigns.GetCampaignDetailsParams{}

	ct := hsClient.Crm().Campaigns()
	response, err := ct.GetCampaignDetailsWithResponse(context.Background(), campaignGuid, &discountsParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}
	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range *response.JSON200.Properties {
			t.Logf("%+v\n", result)
			t.Log("-----")

			// Assuming Properties is a map of key-value pairs
			if result != nil {
				t.Logf("Property: %s\n", result)
			} else {
				t.Log("No properties found.")
			}
			t.Log("-----")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestSaveCampaigns(t *testing.T) {
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

	// Initialize a variable of type Campaigns
	campaign := campaigns.CreateCampaignJSONRequestBody{
		Properties: map[string]string{
			"hs_name":          "New Campaigns Name 4",
			"hs_notes":         "Campaigns notes",
			"hs_currency_code": "USD",
		},
	}

	ct := hsClient.Crm().Campaigns()
	response, err := ct.CreateCampaignWithResponse(context.Background(), campaign)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		if response.JSON201 == nil {
			t.Fatalf("Response contains no results")
		}

		if response.JSON201.Properties != nil {
			for key, value := range *response.JSON201.Properties {
				t.Logf("Key: %s, Value: %+v\n", key, value)
			}
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestUpdateCampaigns(t *testing.T) {
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

	// Initialize a variable of type Campaigns
	campaign := campaigns.UpdateCampaignJSONRequestBody{
		Properties: map[string]string{
			"hs_name":          "New Campaigns Name 2",
			"hs_notes":         "Campaigns notes Updated",
			"hs_currency_code": "USD",
		},
	}

	campaignGuid := "a2a92f29-aee6-4ac3-a5ad-11b2748a67ad"

	ct := hsClient.Crm().Campaigns()
	response, err := ct.UpdateCampaignWithResponse(context.Background(), campaignGuid, campaign)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil {
			t.Fatalf("Response contains no results")
		}

		if response.JSON200.Properties != nil {
			for key, value := range *response.JSON200.Properties {
				t.Logf("Key: %s, Value: %+v\n", key, value)
			}
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestDeleteCampaign(t *testing.T) {
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

	campaignGuid := "a2a92f29-aee6-4ac3-a5ad-11b2748a67ad"

	ct := hsClient.Crm().Campaigns()

	response, err := ct.DeleteCampaignWithResponse(context.Background(), campaignGuid)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Campaign Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
