package association

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/association/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// SchemasDiscovery is the struct that contains all API clients
type SchemasDiscovery struct {
	Schemas *schemas.ClientWithResponses
}

// NewSchemasDiscovery creates a new instance of SchemasDiscovery
func NewSchemasDiscovery(config *configuration.Configuration) (*SchemasDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	objectClient, err := schemas.NewClientWithResponses(config.BasePath, schemas.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	return &SchemasDiscovery{
		Schemas: objectClient,
	}, nil
}
