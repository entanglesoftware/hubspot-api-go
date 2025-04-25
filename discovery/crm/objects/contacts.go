package objects

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
)

// ContactsDiscovery is the struct that contains all API clients
type ContactsDiscovery struct {
	Contacts *contacts.ClientWithResponses
}

// NewContactsDiscovery creates a new instance of ContactsDiscovery
func NewContactsDiscovery(config *configuration.Configuration) (*ContactsDiscovery, error) {
	token, err := config.GetToken()
	fmt.Errorf(os.Getenv("HS_ACCESS_TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	contactClient, err := contacts.NewClientWithResponses(config.BasePath, contacts.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
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
