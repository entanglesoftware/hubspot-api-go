package oauth

// AccessTokenInfoResponse represents the structure of the response for an access token info.
type AccessTokenInfoResponse struct {
	HubID     int      `json:"hubId"`
	UserID    int      `json:"userId"`
	Scopes    []string `json:"scopes"`
	TokenType string   `json:"tokenType"`
	User      *string  `json:"user,omitempty"`
	HubDomain *string  `json:"hubDomain,omitempty"`
	AppID     int      `json:"appId"`
	ExpiresIn int      `json:"expiresIn"`
	Token     string   `json:"token"`
}

// AccessTokenAttributeTypeMap is a static field that holds the attribute information.
var AccessTokenAttributeTypeMap = []struct {
	Name     string `json:"name"`
	BaseName string `json:"baseName"`
	Type     string `json:"type"`
	Format   string `json:"format"`
}{
	{"HubID", "hubId", "int", "int32"},
	{"UserID", "userId", "int", "int32"},
	{"Scopes", "scopes", "Array", "string"},
	{"TokenType", "tokenType", "string", "string"},
	{"User", "user", "string", "string"},
	{"HubDomain", "hubDomain", "string", "string"},
	{"AppID", "appId", "int", "int32"},
	{"ExpiresIn", "expiresIn", "int", "int32"},
	{"Token", "token", "string", "string"},
}

// GetAttributeTypeMap returns the attribute type map for the struct.
func (r *AccessTokenInfoResponse) GetAttributeTypeMap() []struct {
	Name     string `json:"name"`
	BaseName string `json:"baseName"`
	Type     string `json:"type"`
	Format   string `json:"format"`
} {
	return AccessTokenAttributeTypeMap
}

// NewAccessTokenInfoResponse Constructor function to initialize
func NewAccessTokenInfoResponse() *AccessTokenInfoResponse {
	return &AccessTokenInfoResponse{}
}

type AccessTokensApi struct {
	// Define fields and methods for RefreshTokensApi
}
