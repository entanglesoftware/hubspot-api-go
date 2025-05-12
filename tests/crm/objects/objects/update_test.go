package objects_test

import (
	"context"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"
)

func TestUpdateObject(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Make the API call
	objectParams := objects.UpdateObjectParams{}

	ct := crmClient.Objects()

	objectType := "contacts"

	objectId := "100260047027"

	body := objects.UpdateObjectJSONRequestBody{
		Properties: map[string]string{
			"firstname": "John",
			"lastname":  "Doe",
			"email":     "johndoe57hy@example.com",
		},
	}

	response, err := ct.UpdateObjectWithResponse(context.Background(), objectType, objectId, &objectParams, body)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			fmt.Println("No data found.")
			return
		}

		properties := response.JSON200.Properties
		if properties == nil {
			fmt.Println("No properties found.")
			return
		}
		for key, value := range properties {
			fmt.Printf("%s: %s\n", key, value)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
