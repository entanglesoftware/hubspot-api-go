package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/products"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// ProductsDiscovery is the struct that contains all API clients
type ProductsDiscovery struct {
	Products *products.ClientWithResponses
}

// NewProductsDiscovery creates a new instance of ProductsDiscovery
func NewProductsDiscovery(config *configuration.Configuration) (*ProductsDiscovery, error) {
	// Create configuration for API clients
	productClient, err := products.NewClientWithResponses(config.BasePath, products.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &ProductsDiscovery{
		Products: productClient,
	}, nil
}
