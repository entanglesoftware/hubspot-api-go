package products_test

import (
	"context"
	"github.com/entanglesoftware/hubspot-api-go/tests/testsutil"
	"testing"
)

// TestDeleteProductById fetches a page of products
func TestDeleteProductById(t *testing.T) {
	crmClient := testsutil.GetClient()

	ct := crmClient.Products()

	response, err := ct.DeleteProductByIdWithResponse(context.Background(), "17897571956")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("Product Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
