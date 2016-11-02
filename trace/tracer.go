package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
	// ...interface{} means it can take anytype/number of arguments
}

type tracer struct {
	out io.Writer
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
