package crm

import (
	"sync"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/deals"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm/objects"

	"github.com/entanglesoftware/hubspot-api-go/configuration"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/companies"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/leads"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/lineItems"
	customObject "github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/products"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/tickets"
)

// CrmDiscovery provides access to API clients for various CRM modules.
type CrmDiscovery struct {
	config         *configuration.Configuration
	clientRegistry map[string]interface{}
	mu             sync.RWMutex
}

// NewCrmDiscovery initializes and returns a new CrmDiscovery instance.
func NewCrmDiscovery(config *configuration.Configuration) *CrmDiscovery {
	return &CrmDiscovery{
		config:         config,
		clientRegistry: make(map[string]interface{}),
	}
}

// getClient fetches or initializes an API client.
// Ensures thread-safe initialization and caching.
func (d *CrmDiscovery) getClient(key string, constructor func(config *configuration.Configuration) interface{}) interface{} {
	d.mu.RLock()
	client, exists := d.clientRegistry[key]
	d.mu.RUnlock()

	if exists {
		return client
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	// Double-check after acquiring the lock
	client, exists = d.clientRegistry[key]
	if exists {
		return client
	}

	// Initialize and cache the client
	client = constructor(d.config)
	d.clientRegistry[key] = client
	return client
}

// Contacts retrieves the ContactsDiscovery client.
func (d *CrmDiscovery) Contacts() *contacts.ClientWithResponses {
	return d.getClient("contacts", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewContactsDiscovery(config)
		return client.Contacts
	}).(*contacts.ClientWithResponses)
}

// Products retrieves the ProductsDiscovery client.
func (d *CrmDiscovery) Products() *products.ClientWithResponses {
	return d.getClient("products", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewProductsDiscovery(config)
		return client.Products
	}).(*products.ClientWithResponses)
}

// Products retrieves the ProductsDiscovery client.
func (d *CrmDiscovery) Companies() *companies.ClientWithResponses {
	return d.getClient("companies", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewCompaniesDiscovery(config)
		return client.Companies
	}).(*companies.ClientWithResponses)
}

// Deals retrieves the DealsDiscovery client.
func (d *CrmDiscovery) Deals() *deals.ClientWithResponses {
	return d.getClient("deals", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewDealsDiscovery(config)
		return client.Deals
	}).(*deals.ClientWithResponses)
}

// Tickets retrieves the TicketsDiscovery client.
func (d *CrmDiscovery) Tickets() *tickets.ClientWithResponses {
	return d.getClient("tickets", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewTicketsDiscovery(config)
		return client.Tickets
	}).(*tickets.ClientWithResponses)
}

// Leads retrieves the LeadsDiscovery client.
func (d *CrmDiscovery) Leads() *leads.ClientWithResponses {
	return d.getClient("leads", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewLeadsDiscovery(config)
		return client.Leads
	}).(*leads.ClientWithResponses)
}

// LineItems retrieves the LineItemsDiscovery client.
func (d *CrmDiscovery) LineItems() *lineItems.ClientWithResponses {
	return d.getClient("lineItems", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewLineItemsDiscovery(config)
		return client.LineItems
	}).(*lineItems.ClientWithResponses)
}

// SchemaItems retrieves the SchemaItemsDiscovery client.
func (d *CrmDiscovery) SchemaItems() *schemas.ClientWithResponses {
	return d.getClient("schemas", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewSchemaItemsDiscovery(config)
		return client.Schemas
	}).(*schemas.ClientWithResponses)
}

// Objects retrieves the ObjectsDiscovery client.
func (d *CrmDiscovery) Objects() *customObject.ClientWithResponses {
	return d.getClient("objects", func(config *configuration.Configuration) interface{} {
		client, _ := objects.NewObjectsDiscovery(config)
		return client.Objects
	}).(*customObject.ClientWithResponses)
}
