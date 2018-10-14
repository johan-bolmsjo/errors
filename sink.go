package errors

// Sink is an error sink that only remembers the first error sent to it. It may
// be used to simplify some code that can continue executing while saving
// the first error for later.
type Sink struct {
	err error
}

// Cause returns the error stored in the sink.
func (sink *Sink) Cause() error {
	return sink.err
}

// Ok returns true if there is no error in the sink.
func (sink *Sink) Ok() bool {
	return sink.err == nil
}

// Send error to sink. The error is dropped if there is already an error in the sink.
// Returns the sink itself to allow method chaining.
func (sink *Sink) Send(err error) *Sink {
	if sink.err == nil {
		sink.err = err
	}
	return sink
}
