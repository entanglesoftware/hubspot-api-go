package oauth

import (
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"os"
)

// OauthDiscovery holds instances of the API clients
type OauthDiscovery struct {
	accessTokensApi  *AccessTokensApi
	refreshTokensApi *RefreshTokensApi
	tokensApi        *TokensApi
}

// NewOauthDiscovery initializes the OauthDiscovery struct with optional configuration
func NewOauthDiscovery(config configuration.Configuration) *OauthDiscovery {
	return &OauthDiscovery{
		accessTokensApi:  &AccessTokensApi{},
		refreshTokensApi: &RefreshTokensApi{},
		tokensApi:        &TokensApi{},
	}
}

// GetAuthorizationUrl generates the authorization URL based on the provided parameters
func (o *OauthDiscovery) GetAuthorizationUrl(optionalScope, state string) string {
	clientId := os.Getenv("HS_CLIENT_ID")
	clientSecret := os.Getenv("HS_CLIENT_SECRET")
	redirectUri := os.Getenv("HS_REDIRECT_URI")
	scope := os.Getenv("HS_SCOPE")

	if clientId == "" || clientSecret == "" || redirectUri == "" || scope == "" {
		panic("Error: Missing required environment variables (CLIENT_ID, CLIENT_SECRET, REDIRECT_URI, SCOPE).")
		return ""
	}

	url := fmt.Sprintf("https://app.hubspot.com/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s", clientId, redirectUri, scope)

	if optionalScope != "" {
		url += fmt.Sprintf(" %s", optionalScope)
	}

	if state != "" {
		url += fmt.Sprintf("&state=%s", state)
	}

	return url
}
