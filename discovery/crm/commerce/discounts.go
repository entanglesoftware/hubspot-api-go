package commerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/discounts"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// DiscountsDiscovery is the struct that contains all API clients
type DiscountsDiscovery struct {
	Discounts *discounts.ClientWithResponses
}

// NewDiscountsDiscovery creates a new instance of DiscountsDiscovery
func NewDiscountsDiscovery(config *configuration.Configuration) (*DiscountsDiscovery, error) {
	// Create configuration for API clients
	objectClient, err := discounts.NewClientWithResponses(config.BasePath, discounts.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &DiscountsDiscovery{
		Discounts: objectClient,
	}, nil
}
