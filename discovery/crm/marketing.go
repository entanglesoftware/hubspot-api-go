package crm

import (
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/marketing/campaigns"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm/marketing"
)

// Campaigns retrieves the CampaignsDiscovery client.
func (d *CrmDiscovery) Campaigns() *campaigns.ClientWithResponses {
	return d.getClient("campaigns", func(config *configuration.Configuration) interface{} {
		client, _ := marketing.NewCampaignsDiscovery(config)
		return client.Campaigns
	}).(*campaigns.ClientWithResponses)
}
