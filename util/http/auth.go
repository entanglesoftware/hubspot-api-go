package http

type AuthMethod string

const (
	HapiKey     AuthMethod = "hapikey"
	AccessToken AuthMethod = "accessToken"
)
