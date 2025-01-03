package products_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/products"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

func TestUpdateProduct(t *testing.T) {
	// Fetch the access token from the environment
	token := os.Getenv("HS_ACCESS_TOKEN")

	if token == "" {
		t.Skip("HS_ACCESS_TOKEN is not set. Skipping test.")
	}

	// Correctly initialize the struct with the proper syntax
	config := configuration.Configuration{
		AccessToken:            token,
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}

	hsClient := hubspot.NewClient(config)

	// Initialize the client
	hsClient.SetAccessToken(token)

	// Initialize a variable of type Product
	product := products.CreateProductJSONBody{
		Properties: map[string]string{
			"name":   "New Product",
			"price":  "234",
			"hs_sku": "pn",
		},
	}

	productId := "17897571956"

	// Serialize the product properties to JSON
	body, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Error serializing product properties: %v", err)
	}

	contentType := "application/json"

	ct := hsClient.Crm().Products()

	response, err := ct.UpdateProductWithBodyWithResponse(context.Background(), productId, contentType, bytes.NewReader(body))
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
