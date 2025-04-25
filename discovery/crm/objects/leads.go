package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/leads"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// LeadsDiscovery is the struct that contains all API clients
type LeadsDiscovery struct {
	Leads *leads.ClientWithResponses
}

// NewLeadsDiscovery creates a new instance of LeadsDiscovery
func NewLeadsDiscovery(config *configuration.Configuration) (*LeadsDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	ticketClient, err := leads.NewClientWithResponses(config.BasePath, leads.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &LeadsDiscovery{
		Leads: ticketClient,
	}, nil
}
