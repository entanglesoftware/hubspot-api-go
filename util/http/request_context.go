package http

import (
	"net/http"
)

// IRequestContext defines the interface for setting header parameters and agents
type IRequestContext interface {
	SetHeaderParam(key, value string)
	SetAgent(agent *http.Transport)
}

// RequestContext implements the IRequestContext interface
type RequestContext struct {
	headers map[string]string
	agent   *http.Transport
}

// NewRequestContext creates a new instance of RequestContext
func NewRequestContext() *RequestContext {
	return &RequestContext{
		headers: make(map[string]string),
	}
}
