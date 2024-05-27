package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tests-repo/internal/domain"
)

type (
	// UserRepository is a repository for user.
	UserRepository interface {
		InsertOne(ctx context.Context, user domain.User) error
		FindOne(ctx context.Context, id uuid.UUID) (domain.User, error)
		DeleteOne(ctx context.Context, id uuid.UUID) error
	}

	// User is a service for user.
	User struct {
		repo UserRepository
	}
)

// NewUser creates a new user service.
func NewUser(repo UserRepository) *User {
	return &User{
		repo: repo,
	}
}

// Create creates a user.
func (u *User) Create(ctx context.Context, email string) (domain.User, error) {
	user := domain.NewUser(email)

	err := u.repo.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// Get gets a user.
func (u *User) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user, err := u.repo.FindOne(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// Delete deletes a user.
func (u *User) Delete(ctx context.Context, id uuid.UUID) error {
	err := u.repo.DeleteOne(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
