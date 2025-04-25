package products_test

import (
	"context"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

// TestDeleteProductById fetches a page of products
func TestDeleteProductById(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	ct := crm.Products()

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
