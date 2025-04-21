package crm

import (
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/association/details"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/association/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm/association"
)

// Details retrieves the DetailsDiscovery client.
func (d *CrmDiscovery) Details() *details.ClientWithResponses {
	return d.getClient("details", func(config *configuration.Configuration) interface{} {
		client, _ := association.NewDetailsDiscovery(config)
		return client.Details
	}).(*details.ClientWithResponses)
}

// Schemas retrieve the SchemasDiscovery client.
func (d *CrmDiscovery) Schemas() *schemas.ClientWithResponses {
	return d.getClient("schemas", func(config *configuration.Configuration) interface{} {
		client, _ := association.NewSchemasDiscovery(config)
		return client.Schemas
	}).(*schemas.ClientWithResponses)
}
