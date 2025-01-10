package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/tickets"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// TicketsDiscovery is the struct that contains all API clients
type TicketsDiscovery struct {
	Tickets *tickets.ClientWithResponses
}

// NewTicketsDiscovery creates a new instance of TicketsDiscovery
func NewTicketsDiscovery(config *configuration.Configuration) (*TicketsDiscovery, error) {
	// Create configuration for API clients
	ticketClient, err := tickets.NewClientWithResponses(config.BasePath, tickets.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &TicketsDiscovery{
		Tickets: ticketClient,
	}, nil
}
