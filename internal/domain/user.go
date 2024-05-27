package domain

import (
	"github.com/google/uuid"
)

// User is a user.
type User struct {
	// ID is the user ID.
	ID uuid.UUID
	// Email is the user email.
	Email string
}

// NewUser creates a new user.
func NewUser(email string) User {
	return User{
		ID:    uuid.New(),
		Email: email,
	}
}
