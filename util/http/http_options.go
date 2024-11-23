package http

// Options IHttpOptions struct for HTTP options
type Options struct {
	Method      string
	Path        string
	OverlapUrl  string
	QS          map[string]string
	Headers     Headers
	Body        interface{}
	DefaultJSON bool
}
