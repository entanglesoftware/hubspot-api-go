package hubspot

import (
	"errors"
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	decorator "github.com/entanglesoftware/hubspot-api-go/util/decorator"
	"github.com/entanglesoftware/hubspot-api-go/util/http"
	"sync"
	"time"
)

type Client struct {
	Config configuration.Configuration
	// automation               *AutomationDiscovery
	// cms                      *CmsDiscovery
	// communicationPreferences *CommunicationPreferencesDiscovery
	// conversations            *ConversationsDiscovery
	crm *CrmDiscovery
	// events                   *EventsDiscovery
	// files                    *FilesDiscovery
	// marketing                *MarketingDiscovery
	// oauth                    *OauthDiscovery
	// settings                 *SettingsDiscovery
	// webhooks                 *WebhooksDiscovery
	mutex sync.Mutex
}

func NewClient(config configuration.Configuration) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) init() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Initialize or reset service instances
	// c.automation = nil
	// c.cms = nil
	// c.communicationPreferences = nil
	// c.conversations = nil
	c.crm = nil
	// c.events = nil
	// c.files = nil
	// c.marketing = nil
	// c.oauth = nil
	// c.settings = nil
	// c.webhooks = nil
}

// Automation Lazy loading for AutomationDiscovery
// func (c *Client) Automation() *AutomationDiscovery {
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()
//
// 	if c.automation == nil {
// 		c.automation = NewAutomationDiscovery(c.Config)
// 	}
// 	return c.automation
// }
//
// // Cms Lazy loading for CmsDiscovery
// func (c *Client) Cms() *CmsDiscovery {
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()
//
// 	if c.cms == nil {
// 		c.cms = NewCmsDiscovery(c.Config)
// 	}
// 	return c.cms
// }

// SetAccessToken Example implementation for setting access token
func (c *Client) SetAccessToken(token string) {
	c.Config.AccessToken = token
	c.init()
}

// SetApiKey Example implementation for setting API key
func (c *Client) SetApiKey(apiKey string) {
	c.Config.APIKey = apiKey
	c.init()
}

// ApiRequest Making API request with decorators applied
func (c *Client) ApiRequest(opts http.Options) (interface{}, error) {
	request, err := http.NewHttpRequest(c.Config, opts)
	if err != nil {
		fmt.Println("Something went wrong while configuring http")
		return nil, err
	}

	client := http.NewHttpClient()
	send := client.SendAdapter

	// Apply decorators
	decorators := c.getDecorators()
	for _, decor := range decorators {
		send = decor.Decorate(send)
	}

	// Execute the request and return the response
	return send(request)
}

func (c *Client) getDecorators() []decorator.IDecorator {
	var decorators []decorator.IDecorator

	if c.Config.LimiterOptions != nil {
		decorators = append(decorators, decorator.NewLimiterDecorator(time.Second*1000, 5))
	}

	if c.Config.NumberOfAPICallRetries > 0 {
		if c.Config.NumberOfAPICallRetries > 6 {
			panic(errors.New("numberOfApiCallRetries can be set to a number from 0 - 6"))
		}
		decorators = append(decorators, decorator.NewRetryDecorator(c.Config.NumberOfAPICallRetries))
	}

	return decorators
}
