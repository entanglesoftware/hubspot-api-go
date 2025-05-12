package contacts_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
)

func TestUpdateContact(t *testing.T) {
	crmClient := testsutil.GetClient()

	contactId := "87969316513"

	ct := crmClient.Contacts()

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
