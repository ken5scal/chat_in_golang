package trace

import "io"

type Tracer interface {
	Trace(...interface{})
	// ...interface{} means it can take anytype/number of arguments
}

func New(w io.Writer) Tracer {
	return nil
}
