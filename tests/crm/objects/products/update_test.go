package products_test

import (
	"context"
	_ "github.com/entanglesoftware/hubspot-api-go/tests"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/products"
)

func TestUpdateProduct(t *testing.T) {
	crmClient := testsutil.GetClient()

	// Initialize a variable of type Product
	name := "New Product 1"
	price := "234"
	hs_sku := "sku"

	productId := "20156772701"

	ct := crmClient.Products()

	body := products.UpdateProductJSONRequestBody{
		Properties: map[string]string{
			"hs_sku": hs_sku,
			"name":   name,
			"price":  price,
		},
	}

	response, err := ct.UpdateProductWithResponse(context.Background(), productId, body)
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
