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
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	ticketClient, err := schemas.NewClientWithResponses(config.BasePath, schemas.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	return &SchemasDiscovery{
		Schemas: ticketClient,
	}, nil
}
