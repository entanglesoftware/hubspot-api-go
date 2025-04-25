package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/lineItems"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// LineItemsDiscovery is the struct that contains all API clients
type LineItemsDiscovery struct {
	LineItems *lineItems.ClientWithResponses
}

// NewLineItemsDiscovery creates a new instance of LineItemsDiscovery
func NewLineItemsDiscovery(config *configuration.Configuration) (*LineItemsDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	ticketClient, err := lineItems.NewClientWithResponses(config.BasePath, lineItems.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &LineItemsDiscovery{
		LineItems: ticketClient,
	}, nil
}
