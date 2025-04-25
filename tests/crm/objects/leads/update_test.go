package leads_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/leads"
	"github.com/entanglesoftware/hubspot-api-go/configuration"

	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

func TestUpdateLead(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	// Initialize a variable of type Lead
	hs_lead_name := "Update Lead New"
	hs_pipeline_stage := "attempting-stage-id"

	leadId := "396777411872"

	ct := crm.Leads()

	body := leads.UpdateLeadJSONRequestBody{
		Properties: map[string]string{
			"hs_lead_name":      hs_lead_name,
			"hs_pipeline_stage": hs_pipeline_stage,
		},
	}

	response, err := ct.UpdateLeadWithResponse(context.Background(), leadId, body)
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
