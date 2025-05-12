package objects_test

import (
	"context"
	"fmt"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"
)

func TestGetObjects(t *testing.T) {
	crmClient := testsutil.GetClient()

	limit := 10

	// Make the API call
	objectParams := objects.GetObjectsParams{
		Limit: &limit,
	}

	ct := crmClient.Objects()

	objectType := "contacts"

	response, err := ct.GetObjectsWithResponse(context.Background(), objectType, &objectParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range *response.JSON200 {
			t.Log("-----\n")
			t.Logf("%+v\n", result)
			t.Log("-----\n")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestGetObject(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Make the API call
	objectParams := objects.GetObjectByTypeAndIdParams{}

	ct := crmClient.Objects()

	objectType := "contacts"

	objectId := "100260047027"

	response, err := ct.GetObjectByTypeAndIdWithResponse(context.Background(), objectType, objectId, &objectParams)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Id == "" {
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
