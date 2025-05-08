package objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/users"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// UsersDiscovery is the struct that contains all API clients
type UsersDiscovery struct {
	Users *users.ClientWithResponses
}

// NewUsersDiscovery creates a new instance of UsersDiscovery
func NewUsersDiscovery(config *configuration.Configuration) (*UsersDiscovery, error) {
	token, err := config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	userClient, err := users.NewClientWithResponses(config.BasePath, users.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &UsersDiscovery{
		Users: userClient,
	}, nil
}
