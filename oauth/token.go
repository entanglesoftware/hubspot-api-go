package oauth

// TokenResponseIF represents the structure of the token response information.
type TokenResponseIF struct {
	AccessToken  string  `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
	IDToken      *string `json:"idToken,omitempty"` // Optional field
	TokenType    string  `json:"tokenType"`
	ExpiresIn    int     `json:"expiresIn"`
}

// TokenAttributeTypeMap is a static field that holds the attribute information for this struct.
var TokenAttributeTypeMap = []struct {
	Name     string `json:"name"`
	BaseName string `json:"baseName"`
	Type     string `json:"type"`
	Format   string `json:"format"`
}{
	{"AccessToken", "accessToken", "string", "string"},
	{"RefreshToken", "refreshToken", "string", "string"},
	{"IDToken", "idToken", "string", "string"},
	{"TokenType", "tokenType", "string", "string"},
	{"ExpiresIn", "expiresIn", "int", "int32"},
}

// GetAttributeTypeMap returns the attribute type map for the struct.
func (t *TokenResponseIF) GetAttributeTypeMap() []struct {
	Name     string `json:"name"`
	BaseName string `json:"baseName"`
	Type     string `json:"type"`
	Format   string `json:"format"`
} {
	return TokenAttributeTypeMap
}

// NewTokenResponseIF Constructor function to initialize TokenResponseIF
func NewTokenResponseIF() *TokenResponseIF {
	return &TokenResponseIF{}
}

type TokensApi struct {
	// Define fields and methods for TokensApi
}
