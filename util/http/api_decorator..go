package http

import (
	"reflect"

	"github.com/entanglesoftware/hubspot-api-go/util/decorator"
)

// ApiDecoratorService is a singleton service for managing and applying decorators
type ApiDecoratorService struct {
	decorators []decorator.IDecorator
}

var instance *ApiDecoratorService

// GetInstance returns the singleton instance of ApiDecoratorService
func GetInstance() *ApiDecoratorService {
	if instance == nil {
		instance = &ApiDecoratorService{
			decorators: []decorator.IDecorator{},
		}
	}
	return instance
}

// SetDecorators sets the decorators for the service
func (service *ApiDecoratorService) SetDecorators(decorators []decorator.IDecorator) {
	service.decorators = decorators
}

// Apply applies decorators to all methods of an API
func (service *ApiDecoratorService) Apply(api interface{}) interface{} {
	apiValue := reflect.ValueOf(api)
	apiType := apiValue.Type()

	if len(service.decorators) > 0 && apiType.Kind() == reflect.Ptr {
		clientValue := apiValue.Elem()
		clientType := clientValue.Type()

		for i := 0; i < clientType.NumMethod(); i++ {
			method := clientValue.Method(i)
			methodType := method.Type()

			// Skip non-function methods
			if methodType.Kind() != reflect.Func {
				continue
			}

			originalMethod := method.Interface()

			// Check if the method matches the expected signature
			decoratedMethod, ok := originalMethod.(func(args ...interface{}) (interface{}, error))
			if !ok {
				continue // Skip methods that don't match the signature
			}

			// Apply decorators to the method
			for _, svcDecorator := range service.decorators {
				decoratedMethod = svcDecorator.Decorate(decoratedMethod)
			}

			// Use reflection to set the decorated method back to the API
			clientValue.FieldByName(clientType.Method(i).Name).Set(reflect.ValueOf(decoratedMethod))
		}
	}

	return api
}

// ApplyToMethod applies decorators to a single method
func (service *ApiDecoratorService) ApplyToMethod(method func(args ...interface{}) (interface{}, error)) func(args ...interface{}) (interface{}, error) {
	for _, svcDecorator := range service.decorators {
		method = svcDecorator.Decorate(method)
	}
	return method
}
