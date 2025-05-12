package contacts_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
)

func TestMergeContacts(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Contacts()

	body := contacts.MergeContactsJSONRequestBody{
		ObjectIdToMerge: "87484939431",
		PrimaryObjectId: "83910492845",
	}

	response, err := ct.MergeContactsWithResponse(context.Background(), body)
	if err != nil {
		t.Fatalf("API call failed: %v", response)
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
