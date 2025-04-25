package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/companies"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// CompaniesDiscovery is the struct that contains all API clients
type CompaniesDiscovery struct {
	Companies *companies.ClientWithResponses
}

// NewCompaniesDiscovery creates a new instance of CompaniesDiscovery
func NewCompaniesDiscovery(config *configuration.Configuration) (*CompaniesDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	productClient, err := companies.NewClientWithResponses(config.BasePath, companies.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &CompaniesDiscovery{
		Companies: productClient,
	}, nil
}
