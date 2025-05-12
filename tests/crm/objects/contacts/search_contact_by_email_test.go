package contacts_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
)

func TestSearchContactByEmail(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Make the API call

	propertyName := "email"
	email := "johndoe12@example.com"
	limit := 10

	body := contacts.SearchContactsJSONRequestBody{
		Limit: &limit,
		FilterGroups: []contacts.FilterGroups{{
			Filters: []contacts.Filter{{
				Operator:     contacts.FilterOperator("EQ"),
				PropertyName: propertyName,
				Value:        email,
			}},
		}},
	}

	ct := crmClient.Contacts()

	response, err := ct.SearchContactsWithResponse(context.Background(), body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() != 200 {
		t.Fatalf("API call failed with code: %v", response.StatusCode())
	}

	var result struct {
		Total int `json:"total"`
	}

	if err := json.Unmarshal(response.Body, &result); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if result.Total == 0 {
		t.Fatalf("Response contains no results")
	}

	if result.Total > 0 && response.JSON200 != nil && response.JSON200.Results != nil {
		fmt.Printf("Total Result : %+v\n", result.Total)
		conatct := response.JSON200.Results[0]
		fmt.Printf("%+v\n", conatct)
	} else {
		fmt.Printf("Email %s does not exist in HubSpot.\n", email)
	}
}
