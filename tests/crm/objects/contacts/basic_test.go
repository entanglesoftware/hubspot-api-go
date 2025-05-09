package contacts_test

import (
	"context"
	"testing"

	_ "github.com/entanglesoftware/hubspot-api-go/tests"

	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
)

// TestGetContacts fetches a page of contacts
func TestGetContacts(t *testing.T) {
	crm := testsutil.GetClient()

	limit := 10

	// Make the API call
	contactsParams := contacts.GetContactsParams{
		Limit: &limit,
	}

	ct := crm.Contacts()

	response, err := ct.GetContactsWithResponse(context.Background(), &contactsParams)
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

// TestGetContactById fetches a page of contacts
func TestGetContactById(t *testing.T) {
	crm := testsutil.GetClient()

	// Make the API call
	contactByIdParam := contacts.GetContactByIdParams{}

	ct := crm.Contacts()

	response, err := ct.GetContactByIdWithResponse(context.Background(), 84952873394, &contactByIdParam)
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

// TestSaveContacts save a contact
func TestSaveContacts(t *testing.T) {
	crm := testsutil.GetClient()

	body := contacts.CreateContactJSONRequestBody{
		Properties: map[string]string{
			"firstname": "John pqr",
			"lastname":  "Doe abc",
			"email":     "johndoe12111@example.com",
		},
	}

	ct := crm.Contacts()

	response, err := ct.CreateContactWithResponse(context.Background(), body)
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
