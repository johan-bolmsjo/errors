package errors_test

import (
	"github.com/johan-bolmsjo/errors"
	"testing"
)

func TestSink_Ok(t *testing.T) {
	var sink errors.Sink
	if !sink.Ok() {
		t.Fatalf("sink.Ok() = false; want true")
	}
	if !sink.Send(nil).Ok() {
		t.Fatalf("sink.Send(nil).Ok() = false; want true")
	}
	if sink.Send(errors.New("")).Ok() {
		t.Fatalf("sink.Send(errors.New(...)).Ok() = true; want false")
	}
}

func TestSink_SendAndCause(t *testing.T) {
	var sink errors.Sink
	e1 := errors.New("1")
	e2 := errors.New("2")

	sink.Send(e1)
	sink.Send(e2)

	if e := sink.Cause(); e != e1 {
		t.Fatalf("sink.Cause() = %q; want e1", e)
	}
}
