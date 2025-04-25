package commerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/quotes"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// QuotesDiscovery is the struct that contains all API clients
type QuotesDiscovery struct {
	Quotes *quotes.ClientWithResponses
}

// NewQuotesDiscovery creates a new instance of QuotesDiscovery
func NewQuotesDiscovery(config *configuration.Configuration) (*QuotesDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	objectClient, err := quotes.NewClientWithResponses(config.BasePath, quotes.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &QuotesDiscovery{
		Quotes: objectClient,
	}, nil
}
