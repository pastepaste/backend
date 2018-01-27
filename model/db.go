package model

// DB represents a database.
type DB interface {
	// CreatePaste creates a paste.
	CreatePaste(*api.Paste) error

	// GetPaste gets a paste by its slug.
	GetPaste(string) (*api.Paste, error)
}
