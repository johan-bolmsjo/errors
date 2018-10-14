package errors

// Provide Go standard library errors.New since we shadow the package name.

// New returns an error consisting of the given string.
func New(s string) error {
	return &stringError{s}
}

type stringError struct {
	s string
}

// Error implements the error interface.
func (e *stringError) Error() string {
	return e.s
}
