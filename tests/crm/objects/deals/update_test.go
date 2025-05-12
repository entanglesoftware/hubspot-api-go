package deals_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	_ "github.com/entanglesoftware/hubspot-api-go/tests"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/deals"
)

func TestUpdateDeal(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Initialize a variable of type Deal
	dealstage := "appointmentscheduled"
	amount := "234"
	pipeline := "default"
	dealname := "New Deal Name"

	dealId := "33612177707"

	ct := crmClient.Deals()

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
