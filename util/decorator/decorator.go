package decorator

// IDecorator is the interface for decorators.
type IDecorator interface {
	Decorate(method func(args ...interface{}) (interface{}, error)) func(args ...interface{}) (interface{}, error)
}
