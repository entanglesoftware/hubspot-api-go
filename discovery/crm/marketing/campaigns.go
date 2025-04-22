package marketing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/marketing/campaigns"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// CampaignsDiscovery is the struct that contains all API clients
type CampaignsDiscovery struct {
	Campaigns *campaigns.ClientWithResponses
}

// NewCampaignsDiscovery creates a new instance of CampaignsDiscovery
func NewCampaignsDiscovery(config *configuration.Configuration) (*CampaignsDiscovery, error) {
	// Create configuration for API clients
	objectClient, err := campaigns.NewClientWithResponses(config.BasePath, campaigns.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	return &CampaignsDiscovery{
		Campaigns: objectClient,
	}, nil
}
