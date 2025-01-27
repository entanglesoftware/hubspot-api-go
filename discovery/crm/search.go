package crm

import (
	"context"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/crmsearch"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"net/http"
)

type SearchDiscovery struct {
	Search *crmsearch.ClientWithResponses
}

// Search retrieves the SearchDiscovery client.
func (d *CrmDiscovery) Search() *crmsearch.ClientWithResponses {
	return d.getClient("contacts", func(config *configuration.Configuration) interface{} {
		client, _ := newSearchDiscovery(config)
		return client.Search
	}).(*crmsearch.ClientWithResponses)
}

func newSearchDiscovery(config *configuration.Configuration) (*SearchDiscovery, error) {
	searchClient, err := crmsearch.NewClientWithResponses(config.BasePath, crmsearch.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
		return nil
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create search client: %w", err)
	}

	return &SearchDiscovery{
		Search: searchClient,
	}, nil
}
