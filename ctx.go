package errors

import "fmt"

// ctxError provides additional contextual information to an error by wrapping it.
type ctxError struct {
	msg string
	err error
}

// Error implements the error interface.
func (ctxErr *ctxError) Error() string {
	return ctxErr.msg + ": " + ctxErr.err.Error()
}

// Cause returns the wrapped error.
func (ctxErr *ctxError) Cause() error {
	return ctxErr.err
}

// Wrap an error providing additional contextual information.
// Wrap returns nil if err is nil.
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &ctxError{msg, err}
}

// Wrap an error providing additional contextual information formated according
// to a format specifier. Wrap returns nil if err is nil.
func Wrapf(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return &ctxError{fmt.Sprintf(format, a...), err}
}

// Cause finds the bottom most error in a chain of wrapped errors.
// It does this by calling the method "Cause() error" on err if it exists.
// This is performed recursively until an error is found which does not have the method.
func Cause(err error) error {
	type causer interface{ Cause() error }
	if wrapped, ok := err.(causer); ok {
		return Cause(wrapped.Cause())
	}
	return err
}
