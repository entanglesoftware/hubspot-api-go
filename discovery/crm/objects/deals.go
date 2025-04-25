package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/deals"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// DealsDiscovery is the struct that contains all API clients
type DealsDiscovery struct {
	Deals *deals.ClientWithResponses
}

// NewDealsDiscovery creates a new instance of DealsDiscovery
func NewDealsDiscovery(config *configuration.Configuration) (*DealsDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	dealClient, err := deals.NewClientWithResponses(config.BasePath, deals.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &DealsDiscovery{
		Deals: dealClient,
	}, nil
}
