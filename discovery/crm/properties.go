package crm

import (
	"context"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/properties"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"net/http"
)

// PropertiesDiscovery is the struct that contains all API clients
type PropertiesDiscovery struct {
	Properties *properties.ClientWithResponses
}

// NewPropertiesDiscovery creates a new instance of PropertiesDiscovery
func NewPropertiesDiscovery(config *configuration.Configuration) (*PropertiesDiscovery, error) {
	// Create configuration for API clients
	propertiesClient, err := properties.NewClientWithResponses(config.BasePath, properties.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	return &PropertiesDiscovery{
		Properties: propertiesClient,
	}, nil
}

// Properties retrieve the PropertiesDiscovery client.
func (d *CrmDiscovery) Properties() *properties.ClientWithResponses {
	return d.getClient("properties", func(config *configuration.Configuration) interface{} {
		client, _ := NewPropertiesDiscovery(config)
		return client.Properties
	}).(*properties.ClientWithResponses)
}
