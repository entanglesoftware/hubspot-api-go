package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// SchemasDiscovery is the struct that contains all API clients
type SchemasDiscovery struct {
	Schemas *schemas.ClientWithResponses
}

// NewSchemaItemsDiscovery creates a new instance of SchemasDiscovery
func NewSchemaItemsDiscovery(config *configuration.Configuration) (*SchemasDiscovery, error) {
	// Create configuration for API clients
	ticketClient, err := schemas.NewClientWithResponses(config.BasePath, schemas.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	return &SchemasDiscovery{
		Schemas: ticketClient,
	}, nil
}
