package contacts_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestUpdateContact(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)

	contactId := "87969316513"

	ct := hsClient.Crm().Contacts()

	body := contacts.UpdateContactJSONRequestBody{
		Properties: map[string]string{
			"firstname": "John Update",
			"lastname":  "Doe new",
			"email":     "johndoe1211122@example.com",
		},
	}

	response, err := ct.UpdateContactWithResponse(context.Background(), contactId, body)
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
