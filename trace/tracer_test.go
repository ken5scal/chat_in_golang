package trace

import (
	"testing"
	"bytes"
)

// Any method starts from Test and takes one *testing.T type argument is
// recognized as Unit Test
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New returned nil")
	} else {
		tracer.Trace("hello, package trace")
		if buf.String() != "hello, package trace\n" {
			t.Errorf("'%s' is returned as wrong string", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("Data")
}
