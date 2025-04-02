package association

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/association/details"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// DetailsDiscovery is the struct that contains all API clients
type DetailsDiscovery struct {
	Details *details.ClientWithResponses
}

// NewDetailsDiscovery creates a new instance of DetailsDiscovery
func NewDetailsDiscovery(config *configuration.Configuration) (*DetailsDiscovery, error) {
	// Create configuration for API clients
	objectClient, err := details.NewClientWithResponses(config.BasePath, details.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	return &DetailsDiscovery{
		Details: objectClient,
	}, nil
}
