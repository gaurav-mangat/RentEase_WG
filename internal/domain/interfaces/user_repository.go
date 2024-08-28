package interfaces

import (
	"context"
	"rentease/internal/domain/entities"
)

type UserRepo interface {
	SaveUser(user entities.User) error
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	CheckPassword(ctx context.Context, username string, password string) (bool, error)
	UpdateUser(user entities.User) error
	Delete(username string) error
	FindAll() ([]entities.User, error)
}

// Admin
