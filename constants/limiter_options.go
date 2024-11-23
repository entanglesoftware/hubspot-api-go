package constants

// LimiterOptions defines the structure for limiter settings
type LimiterOptions struct {
	MinTime       float64
	MaxConcurrent int
	ID            string
}

// DEFAULT_LIMITER_OPTIONS Default limiter options
var DEFAULT_LIMITER_OPTIONS = LimiterOptions{
	MinTime:       1000.0 / 9,
	MaxConcurrent: 6,
	ID:            "hubspot-client-limiter",
}

// SEARCH_LIMITER_OPTIONS Search limiter options
var SEARCH_LIMITER_OPTIONS = LimiterOptions{
	MinTime:       550.0,
	MaxConcurrent: 3,
	ID:            "search-hubspot-client-limiter",
}
