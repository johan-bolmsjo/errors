package errors_test

import (
	"github.com/johan-bolmsjo/errors"
	"testing"
)

func TestCtx_WrapAndCause(t *testing.T) {
	e1 := errors.Wrap(nil, "")
	if e1 != nil {
		t.Fatalf("errors.Wrap(nil, ...) = %q; want nil", e1)
	}

	orig := errors.New("original")
	wrap1 := errors.Wrap(orig, "aux1")
	wrap2 := errors.Wrap(wrap1, "aux2")

	const want = "aux2: aux1: original"
	wrap2error := wrap2.Error()
	if wrap2error != want {
		t.Fatalf("wrap2.Error() = %q; want %q", wrap2error, want)
	}

	wrap1cause := errors.Cause(wrap1)
	wrap2cause := errors.Cause(wrap2)
	if wrap1cause != wrap2cause {
		t.Fatalf("wrap1cause = %q; want %q", wrap1cause, wrap2cause)
	}
	if wrap2cause != orig {
		t.Fatalf("wrap1cause = %q; want 'orig'", wrap2cause)
	}
}

func TestCtx_Wrapf(t *testing.T) {
	e1 := errors.Wrapf(nil, "")
	if e1 != nil {
		t.Fatalf("errors.Wrapf(nil, ...) = %q; want nil", e1)
	}

	err := errors.Wrapf(errors.New("original"), "foo=%s", "bar")
	got := err.Error()
	const want = "foo=bar: original"
	if got != want {
		t.Fatalf("errors.Wrapwf(...) = %q; want %q", got, want)
	}
}
