package objects

import (
	"context"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"net/http"
)

// ContactsDiscovery is the struct that contains all API clients
type ContactsDiscovery struct {
	Contacts *contacts.Client
}

// NewContactsDiscovery creates a new instance of ContactsDiscovery
func NewContactsDiscovery(config *configuration.Configuration) (*ContactsDiscovery, error) {
	// Create configuration for API clients
	contactClient, err := contacts.NewClient(config.BasePath, contacts.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create configuration: %w", err)
	}

	// Initialize API clients and apply decorators
	// decoratorService := services.GetApiDecoratorServiceInstance()

	return &ContactsDiscovery{
		Contacts: contactClient,
	}, nil
}

// GetAll retrieves all contacts using the BasicApi client
// func (cd *ContactsDiscovery) GetAll(
// 	limit int,
// 	after string,
// 	properties []string,
// 	propertiesWithHistory []string,
// 	associations []string,
// 	archived bool,
// ) ([]api.SimplePublicObjectWithAssociations, error) {
// 	// Use the service to get all contacts
// 	return services.GetAll(cd.BasicApi, limit, after, properties, propertiesWithHistory, associations, archived)
// }
