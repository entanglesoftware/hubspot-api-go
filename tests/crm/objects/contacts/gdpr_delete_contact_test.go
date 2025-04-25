package contacts_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
	"github.com/entanglesoftware/hubspot-api-go/configuration"

	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

func TestGDPRDeleteContact(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	ct := crm.Contacts()

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
