package trace

type Tracer interface {
	Trace(...interface{})
	// ...interface{} means it can take anytype/number of arguments
}
