package associations_test

import (
	"context"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/association/details"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

// TestGetAssociationsDetails fetches a page of details
func TestGetAssociationsDetails(t *testing.T) {
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
	objectType := "contact"
	objectId := "100260047027"
	toObjectType := "companies"

	// Make the API call
	detailsParams := details.GetAssociationsDetailsParams{
		Limit: &limit,
	}

	ct := hsClient.Crm().Details()
	response, err := ct.GetAssociationsDetailsWithResponse(context.Background(), objectType, objectId, toObjectType, &detailsParams)
	if err != nil {
		t.Fatalf("API call failed: %+v", err)
	}
	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range *response.JSON200.Results {
			t.Log("-----")
			t.Logf("To ObjectId: %s\n", result.ToObjectId)
			if result.AssociationTypes != nil {
				for key, value := range result.AssociationTypes {
					t.Logf("Key: %+v, Value: %+v\n", key, value)
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

// TestCreateAssociationsDetails fetches a page of details
func TestCreateAssociationsDetails(t *testing.T) {
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
	objectType := "contact"
	objectId := "100260047027"
	toObjectType := "companies"
	toObjectId := "28292686395"

	// Make the API call
	detailsParams := details.CreateAssociationsDetailsJSONBody{
		{
			AssociationCategory: details.CreateAssociationsDetailsJSONBodyAssociationCategory("HUBSPOT_DEFINED"),
			AssociationTypeId:   1,
		},
	}

	ct := hsClient.Crm().Details()
	response, err := ct.CreateAssociationsDetailsWithResponse(context.Background(), objectType, objectId, toObjectType, toObjectId, detailsParams)
	if err != nil {
		t.Fatalf("API call failed: %+v", err)
	}
	if response.StatusCode() == 201 {
		if response.JSON201 == nil {
			t.Fatalf("Response contains no results")
		}

		if response.JSON201 != nil {
			t.Logf("FromObjectId: %+v\n", response.JSON201.FromObjectId)
			t.Logf("FromObjectTypeId: %+v\n", response.JSON201.FromObjectTypeId)
			t.Logf("ToObjectId: %+v\n", response.JSON201.ToObjectId)
			t.Logf("ToObjectTypeId: %+v\n", response.JSON201.ToObjectTypeId)
			t.Logf("Labels: %+v\n", response.JSON201.Labels)
			if response.JSON201.Labels != nil {
				for _, value := range response.JSON201.Labels {
					t.Logf("Label: %+v\n", value)
				}
			}
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

// TestCreateDefaultAssociationsDetails fetches a page of details
func TestCreateDefaultAssociationsDetails(t *testing.T) {
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
	FromObjectType := "contact"
	FromObjectId := "100260047027"
	toObjectType := "companies"
	toObjectId := "28292686395"

	ct := hsClient.Crm().Details()
	response, err := ct.CreateDefaultAssociationsDetailsWithResponse(context.Background(), FromObjectType, FromObjectId, toObjectType, toObjectId)
	if err != nil {
		t.Fatalf("API call failed: %+v", err)
	}
	if response.StatusCode() == 200 {
		if response.JSON200 == nil {
			t.Fatalf("Response contains no results")
		}

		if response.JSON200 != nil {
			t.Logf("status: %+v\n", response.JSON200.Status)
			t.Logf("startedAt: %+v\n", response.JSON200.StartedAt)
			t.Logf("completedAt: %+v\n", response.JSON200.CompletedAt)
			if response.JSON200.Results != nil {
				for _, value := range response.JSON200.Results {
					t.Logf("Result: %+v\n", value)
				}
			}
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
