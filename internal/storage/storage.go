package storage

import (
	"context"

	"MetroLab/internal/entity"
)

var _ Storage = &PGStorage{}

type Storage interface {
	GetUser(ctx context.Context, userID int64) (*entity.User, error)
}
