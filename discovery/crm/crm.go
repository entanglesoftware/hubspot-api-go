package crm

import (
	"sync"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
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
