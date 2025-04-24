package companies_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/companies"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestUpdateCompany(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)
	domain := "newdomain.com"
	name := "New Name"

	companyId := "28481339557"

	body := companies.UpdateCompanyJSONRequestBody{
		Properties: map[string]string{
			"domain": domain,
			"name":   name,
		},
	}

	ct := hsClient.Crm().Companies()

	response, err := ct.UpdateCompanyWithResponse(context.Background(), companyId, body)
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
