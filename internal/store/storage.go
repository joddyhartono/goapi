package store

import (
	"context"
	"database/sql"

	"github.com/joddyhartono/goapi/internal/models"
)

type Storage struct {
	Users UserRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &userRepository{db: db},
	}
}