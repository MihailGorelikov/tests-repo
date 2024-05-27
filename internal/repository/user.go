package repository

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/tests-repo/internal/domain"
	"github.com/tests-repo/internal/platform/errors"
)

// User is a repository for user.
type User struct {
	mu     sync.Mutex
	values map[uuid.UUID]domain.User
}

// NewUser creates a new user repository.
func NewUser() *User {
	return &User{
		mu:     sync.Mutex{},
		values: make(map[uuid.UUID]domain.User),
	}
}

// InsertOne inserts a user.
func (u *User) InsertOne(_ context.Context, user domain.User) error {
	_, ok := u.values[user.ID]
	if ok {
		return errors.ErrDuplicateRecord
	}

	u.values[user.ID] = user
	return nil
}

// FindOne finds a user by ID.
func (u *User) FindOne(_ context.Context, id uuid.UUID) (domain.User, error) {
	user, ok := u.values[id]
	if !ok {
		return domain.User{}, errors.ErrRecordNotFound
	}

	return user, nil
}

func (u *User) DeleteOne(_ context.Context, id uuid.UUID) (domain.User, error) {
	user, ok := u.values[id]
	if !ok {
		return domain.User{}, errors.ErrRecordNotFound
	}

	delete(u.values, id)
	return user, nil
}
