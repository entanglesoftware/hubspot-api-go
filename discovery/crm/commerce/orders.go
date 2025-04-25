package commerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/orders"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// OrdersDiscovery is the struct that contains all API clients
type OrdersDiscovery struct {
	Orders *orders.ClientWithResponses
}

// NewOrdersDiscovery creates a new instance of OrdersDiscovery
func NewOrdersDiscovery(config *configuration.Configuration) (*OrdersDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	objectClient, err := orders.NewClientWithResponses(config.BasePath, orders.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	return &OrdersDiscovery{
		Orders: objectClient,
	}, nil
}
