package store

import (
	"context"
	"lectures-6/internal/models"
)

type Store interface {
	Create(ctx context.Context, item *models.Item) error
	All(ctx context.Context) ([]*models.Item, error)
	ByID(ctx context.Context, id int) (*models.Item, error)
	Update(ctx context.Context, item *models.Item) error
	Delete(ctx context.Context, id int) error
}

