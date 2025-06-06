package leads_test

import (
	"context"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/leads"

	"github.com/entanglesoftware/hubspot-api-go/hubspot"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

func TestGetLeads(t *testing.T) {
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

	limit := 10

	// Make the API call
	leadParams := leads.GetLeadsParams{
		Limit: &limit,
	}

	ct := hsClient.Crm().Leads()

	response, err := ct.GetLeadsWithResponse(context.Background(), &leadParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range response.JSON200.Results {
			t.Logf("%+v\n", result)
			t.Log("-----")

			// Assuming Properties is a map of key-value pairs
			if result.Properties != nil {
				for key, value := range result.Properties {
					t.Logf("Key: %s, Value: %+v\n", key, value)
				}
			} else {
				t.Log("No properties found.")
			}
			t.Log("-----")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestGetLeadById(t *testing.T) {
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

	// Make the API call
	leadByIdParam := leads.GetLeadByIdParams{}

	ct := hsClient.Crm().Leads()

	response, err := ct.GetLeadByIdWithResponse(context.Background(), "396711567278", &leadByIdParam)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			t.Fatalf("Response contains no results")
		}

		for key, result := range response.JSON200.Properties {
			t.Logf("Key: %s, Value: %+v\n", key, result)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func ptr[T any](v T) *T {
	return &v
}

func TestSaveLead(t *testing.T) {
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
	associationTypeId := int32(580)
	associationCategory := "HUBSPOT_DEFINED"
	id := "28106025611"

	ct := hsClient.Crm().Leads()
	body := leads.CreateLeadJSONRequestBody{
		Properties: map[string]string{
			"hs_lead_name":      "New Lead Demo 1",
			"hs_pipeline_stage": "connected-stage-id",
		},
		Associations: []struct {
			To *struct {
				Id *string "json:\"id,omitempty\""
			} "json:\"to,omitempty\""
			Types *[]struct {
				AssociationCategory *leads.CreateLeadJSONBodyAssociationsTypesAssociationCategory "json:\"associationCategory,omitempty\""
				AssociationTypeId   *int32                                                        "json:\"associationTypeId,omitempty\""
			} "json:\"types,omitempty\""
		}{
			{
				To: &struct {
					Id *string "json:\"id,omitempty\""
				}{
					Id: &id, // Replace with actual ID
				},
				Types: &[]struct {
					AssociationCategory *leads.CreateLeadJSONBodyAssociationsTypesAssociationCategory "json:\"associationCategory,omitempty\""
					AssociationTypeId   *int32                                                        "json:\"associationTypeId,omitempty\""
				}{
					{
						AssociationCategory: (*leads.CreateLeadJSONBodyAssociationsTypesAssociationCategory)(&associationCategory), // Replace with actual category
						AssociationTypeId:   &associationTypeId,                                                                    // Replace with actual association type ID
					},
				},
			},
		},
	}

	response, err := ct.CreateLeadWithResponse(context.Background(), body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		if response.JSON201 == nil || response.JSON201.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON201.Properties != nil {
			t.Logf("Properties: %s", response.JSON201.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
