package dynamodb

// NotFoundError occurs when an item couldn't be found.
type NotFoundError struct {
	// AWSError is the underlying AWS error.
	AWSError error
}

// Error returns the underlying error string.
func (err *NotFoundError) Error() string {
	return err.AWSError.Error()
}
