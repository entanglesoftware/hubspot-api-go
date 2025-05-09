package associations_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/association/schemas"
)

// TestGetAssociationsSchemas fetches a page of schemas
func TestGetAssociationsSchemas(t *testing.T) {
	crmClient := testsutil.GetClient()

	fromObjectType := "contact"
	toObjectType := "companies"

	ct := crmClient.Schemas()
	response, err := ct.GetAssociationsSchemaWithResponse(context.Background(), fromObjectType, toObjectType)
	if err != nil {
		t.Fatalf("API call failed: %+v", err)
	}
	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range response.JSON200.Results {
			t.Log("-----")
			t.Logf("Schemas: %+v\n", result)
			t.Log("-----")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

// TestCreateAssociationsSchemas fetches a page of schemas
func TestCreateAssociationsSchemas(t *testing.T) {
	crmClient := testsutil.GetClient()
	fromObjectType := "contact"
	toObjectType := "companies"

	// Make the API call
	schemasParams := schemas.CreateAssociationSchemaJSONRequestBody{
		InverseLabel: "Custom Schema",
		Name:         "Custom Contact Companies Name New",
		Label:        "Custom Contact Companies Label New",
	}

	ct := crmClient.Schemas()
	response, err := ct.CreateAssociationSchemaWithResponse(context.Background(), fromObjectType, toObjectType, schemasParams)
	if err != nil {
		t.Fatalf("API call failed: %+v", err)
	}
	if response.StatusCode() == 200 {
		if response.JSON200 == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range response.JSON200.Results {
			t.Log("-----")
			t.Logf("Schemas: %+v\n", result)
			t.Log("-----")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

// TestUpdateAssociationsSchemas fetches a page of schemas
func TestUpdateAssociationsSchemas(t *testing.T) {
	crmClient := testsutil.GetClient()
	fromObjectType := "contact"
	toObjectType := "companies"

	schemasParams := schemas.UpdateAssociationSchemaJSONRequestBody{
		InverseLabel:      "Custom Schema Updated",
		AssociationTypeId: 75,
		Label:             "Custom Contact Companies Label Updated",
	}

	ct := crmClient.Schemas()
	response, err := ct.UpdateAssociationSchemaWithResponse(context.Background(), fromObjectType, toObjectType, schemasParams)
	if err != nil {
		t.Fatalf("API call failed: %+v", err)
	}
	if response.StatusCode() == 204 {
		t.Logf("Associations Schema Updated")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
