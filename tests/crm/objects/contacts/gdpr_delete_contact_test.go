package contacts_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
)

func TestGDPRDeleteContact(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Contacts()

	body := contacts.GdprDeleteContactJSONRequestBody{
		IdProperty: "email",
		ObjectId:   "example2@example.com",
	}

	response, err := ct.GdprDeleteContactWithResponse(context.Background(), body)

	if err != nil {
		t.Fatalf("API call failed: %v", response)
	}
	// t.Logf("API call failed: %v", response.JSON204)
	if response.StatusCode() == 204 {
		t.Logf("Contact Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
