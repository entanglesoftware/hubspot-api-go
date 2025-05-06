package crm

import (
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/carts"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/discounts"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/invoices"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/orders"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/quotes"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/commerce/taxes"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm/commerce"
)

// Quotes retrieves the QuotesDiscovery client.
func (d *CrmDiscovery) Quotes() *quotes.ClientWithResponses {
	return d.getClient("quotes", func(config *configuration.Configuration) interface{} {
		client, _ := commerce.NewQuotesDiscovery(config)
		return client.Quotes
	}).(*quotes.ClientWithResponses)
}

// Invoices retrieves the InvoicesDiscovery client.
func (d *CrmDiscovery) Invoices() *invoices.ClientWithResponses {
	return d.getClient("invoices", func(config *configuration.Configuration) interface{} {
		client, _ := commerce.NewInvoicesDiscovery(config)
		return client.Invoices
	}).(*invoices.ClientWithResponses)
}

// Discounts retrieves the DiscountsDiscovery client.
func (d *CrmDiscovery) Discounts() *discounts.ClientWithResponses {
	return d.getClient("discounts", func(config *configuration.Configuration) interface{} {
		client, _ := commerce.NewDiscountsDiscovery(config)
		return client.Discounts
	}).(*discounts.ClientWithResponses)
}

// Taxes retrieves the TaxesDiscovery client.
func (d *CrmDiscovery) Taxes() *taxes.ClientWithResponses {
	return d.getClient("taxes", func(config *configuration.Configuration) interface{} {
		client, _ := commerce.NewTaxesDiscovery(config)
		return client.Taxes
	}).(*taxes.ClientWithResponses)
}

// Orders retrieves the OrdersDiscovery client.
func (d *CrmDiscovery) Orders() *orders.ClientWithResponses {
	return d.getClient("orders", func(config *configuration.Configuration) interface{} {
		client, _ := commerce.NewOrdersDiscovery(config)
		return client.Orders
	}).(*orders.ClientWithResponses)
}

// Carts retrieves the CartsDiscovery client.
func (d *CrmDiscovery) Carts() *carts.ClientWithResponses {
	return d.getClient("carts", func(config *configuration.Configuration) interface{} {
		client, _ := commerce.NewCartsDiscovery(config)
		return client.Carts
	}).(*carts.ClientWithResponses)
}
