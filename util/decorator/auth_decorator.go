package decorator

import (
	"fmt"
	"github.com/entanglesoftware/hubspot-api-go/util/http"
)

type AuthDecorator struct {
	Token string
}

func NewAuthDecorator(token string) *AuthDecorator {
	return &AuthDecorator{
		Token: token,
	}
}

func (d *AuthDecorator) Decorate(method func(args ...interface{}) (interface{}, error)) func(args ...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		if len(args) < 1 {
			return method(args...)
		}

		request, ok := args[0].(*http.Request)
		if !ok {
			return method(args...)
		}
		if request.Headers == nil {
			request.Headers = make(map[string]string)
		}
		if d.Token != "" {
			request.Headers["Authorization"] = fmt.Sprintf("Bearer %s", d.Token)
		}

		return method(args...)
	}
}
