package testsutil

import (
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
	"log"
	"sync"
)

var (
	client *crm.CrmDiscovery
	once   sync.Once
)

func GetClient() *crm.CrmDiscovery {
	once.Do(func() {
		log.Println("Initializing Hubspot client")
		config := configuration.Configuration{
			BasePath:               configuration.BaseURL,
			NumberOfAPICallRetries: 3,
		}
		client = crm.NewCrmDiscovery(&config)
	})
	return client
}
