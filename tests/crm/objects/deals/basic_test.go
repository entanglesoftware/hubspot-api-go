package deals_test

import (
	"context"
	"testing"

	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/deals"
)

func TestGetDeals(t *testing.T) {
	crmClient := testsutil.GetClient()

	limit := 10

	// Make the API call
	dealParams := deals.GetDealsParams{
		Limit: &limit,
	}

	ct := crmClient.Deals()

	response, err := ct.GetDealsWithResponse(context.Background(), &dealParams)
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

func TestGetDealById(t *testing.T) {
	crmClient := testsutil.GetClient()
	// Make the API call
	dealByIdParam := deals.GetDealByIdParams{}

	ct := crmClient.Deals()

	response, err := ct.GetDealByIdWithResponse(context.Background(), "31738621965", &dealByIdParam)
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

func TestSaveDeal(t *testing.T) {
	crmClient := testsutil.GetClient()

	body := deals.CreateDealJSONRequestBody{
		Properties: map[string]string{
			"dealname":  "New Deal 11",
			"dealstage": "appointmentscheduled",
			"pipeline":  "default",
			"amount":    "545",
		},
	}

	ct := crmClient.Deals()

	response, err := ct.CreateDealWithResponse(context.Background(), body)
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
