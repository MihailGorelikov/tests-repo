package errors

import "errors"

var (
	// ErrNotFound is used when a resource is not found.
	ErrRecordNotFound = errors.New("record not found")

	// ErrDuplicateRecord is used when a record already exists.
	ErrDuplicateRecord = errors.New("duplicate record")
)
