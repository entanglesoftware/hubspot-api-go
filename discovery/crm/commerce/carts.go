package commerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/carts"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// CartsDiscovery is the struct that contains all API clients
type CartsDiscovery struct {
	Carts *carts.ClientWithResponses
}

// NewCartsDiscovery creates a new instance of CartsDiscovery
func NewCartsDiscovery(config *configuration.Configuration) (*CartsDiscovery, error) {
	objectClient, err := carts.NewClientWithResponses(config.BasePath, carts.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &CartsDiscovery{
		Carts: objectClient,
	}, nil
}
