package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// ObjectsDiscovery is the struct that contains all API clients
type ObjectsDiscovery struct {
	Objects *objects.ClientWithResponses
}

// NewObjectsDiscovery creates a new instance of ObjectsDiscovery
func NewObjectsDiscovery(config *configuration.Configuration) (*ObjectsDiscovery, error) {
	// Create configuration for API clients
	objectClient, err := objects.NewClientWithResponses(config.BasePath, objects.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &ObjectsDiscovery{
		Objects: objectClient,
	}, nil
}
