package trace

import (
	"testing"
	"bytes"

)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer ==nil {
		t.Error("Tracer is null")
	} else {
		tracer.Trace("Hello trace package")
		if buf.String() != "Hello trace package" {
			t.Errorf("'%s' is not expected", buf.String())
		}
	}
}
