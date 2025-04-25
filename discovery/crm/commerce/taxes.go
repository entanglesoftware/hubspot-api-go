package commerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/taxes"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// TaxesDiscovery is the struct that contains all API clients
type TaxesDiscovery struct {
	Taxes *taxes.ClientWithResponses
}

// NewTaxesDiscovery creates a new instance of TaxesDiscovery
func NewTaxesDiscovery(config *configuration.Configuration) (*TaxesDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	objectClient, err := taxes.NewClientWithResponses(config.BasePath, taxes.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &TaxesDiscovery{
		Taxes: objectClient,
	}, nil
}
