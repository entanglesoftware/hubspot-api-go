package testsutil

import (
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
	"github.com/entanglesoftware/hubspot-api-go/util/decorator"
	"log"
	"os"
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
			AuthDecorator:          decorator.NewAuthDecorator(os.Getenv("HS_ACCESS_TOKEN")),
		}
		client = crm.NewCrmDiscovery(&config)
	})
	return client
}
