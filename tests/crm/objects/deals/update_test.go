package deals_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/deals"
	"github.com/entanglesoftware/hubspot-api-go/configuration"

	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

func TestUpdateDeal(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	// Initialize a variable of type Deal
	dealstage := "appointmentscheduled"
	amount := "234"
	pipeline := "default"
	dealname := "New Deal Name"

	dealId := "33612177707"

	ct := crm.Deals()

	body := deals.UpdateDealJSONRequestBody{
		Properties: map[string]string{
			"pipeline":  pipeline,
			"dealstage": dealstage,
			"amount":    amount,
			"dealname":  dealname,
		},
	}

	response, err := ct.UpdateDealWithResponse(context.Background(), dealId, body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON200.Properties != nil {
			t.Logf("Properties: %s", response.JSON200.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
