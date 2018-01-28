package service

// DB represents a database.
type DB interface {
	// CreatePaste creates a paste.
	CreatePaste(*Paste) error

	// GetPaste gets a paste with the given slug.
	GetPaste(string) (*Paste, error)
}
