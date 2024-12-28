package configuration

import (
	"net/http"
	"time"
)

const (
	BaseURL = "https://api.hubapi.com"
)

// LimiterOptions represents the configuration options for a rate limiter.
type LimiterOptions struct {
	MaxConcurrent            int           // Maximum concurrent jobs
	MinTime                  time.Duration // Minimum time between each job
	Reservoir                int           // Initial tokens in the reservoir
	ReservoirRefreshInterval time.Duration // Interval to refresh reservoir tokens
	ReservoirRefreshAmount   int           // Tokens added on each refresh
}

// LimiterJobOptions represents options for individual jobs within the rate limiter.
type LimiterJobOptions struct {
	Expiration time.Duration // Time after which the job expires
}

// Configuration struct holds all settings required for API requests.
type Configuration struct {
	APIKey                 string             // API key for authentication
	AccessToken            string             // Access token for authentication
	DeveloperAPIKey        string             // Developer API key
	BasePath               string             // Base URL for API requests
	DefaultHeaders         map[string]string  // Default headers for requests
	NumberOfAPICallRetries int                // API call retry count
	LimiterOptions         *LimiterOptions    // Rate limiter configuration
	LimiterJobOptions      *LimiterJobOptions // Per-job rate limiter options
	HTTPAgent              *http.Transport    // HTTP transport agent
}

// NewConfiguration creates a new Configuration.
func NewConfiguration(
	apiKey, accessToken, developerAPIKey, basePath string,
	defaultHeaders map[string]string,
	retries int,
	limiterOptions *LimiterOptions,
	jobOptions *LimiterJobOptions,
	httpAgent *http.Transport,
) *Configuration {
	return &Configuration{
		APIKey:                 apiKey,
		AccessToken:            accessToken,
		DeveloperAPIKey:        developerAPIKey,
		BasePath:               basePath,
		DefaultHeaders:         defaultHeaders,
		NumberOfAPICallRetries: retries,
		LimiterOptions:         limiterOptions,
		LimiterJobOptions:      jobOptions,
		HTTPAgent:              httpAgent,
	}
}
