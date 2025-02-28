package commerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/invoices"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// InvoicesDiscovery is the struct that contains all API clients
type InvoicesDiscovery struct {
	Invoices *invoices.ClientWithResponses
}

// NewInvoicesDiscovery creates a new instance of InvoicesDiscovery
func NewInvoicesDiscovery(config *configuration.Configuration) (*InvoicesDiscovery, error) {
	// Create configuration for API clients
	objectClient, err := invoices.NewClientWithResponses(config.BasePath, invoices.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &InvoicesDiscovery{
		Invoices: objectClient,
	}, nil
}
