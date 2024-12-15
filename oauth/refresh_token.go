package oauth

// RefreshTokenInfoResponse represents the structure of the response for a refresh token info.
type RefreshTokenInfoResponse struct {
	HubID     int      `json:"hubId"`
	UserID    int      `json:"userId"`
	Scopes    []string `json:"scopes"`
	TokenType string   `json:"tokenType"`
	User      *string  `json:"user,omitempty"`      // Optional field
	HubDomain *string  `json:"hubDomain,omitempty"` // Optional field
	ClientID  string   `json:"clientId"`
	Token     string   `json:"token"`
}

type RefreshTokensApi struct {
	// Define fields and methods for RefreshTokensApi
}
