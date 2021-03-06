package errors

// Error represents a handler error. It provides methods for a HTTP status 
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Err  error
	Code int
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

//Status Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}
