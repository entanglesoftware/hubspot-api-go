package http

type AuthType string

const (
	APIKey          AuthType = AuthType(HapiKey)
	AccessTokenKey  AuthType = AuthType(AccessToken)
	DeveloperAPIKey AuthType = AuthType(HapiKey)
)
