package cache

import (
	"context"
	"lectures-6/internal/models"
)

type Cache interface {
	Close() error

	Laptops() LaptopsCacheRepo
	DeleteAll(ctx context.Context) error
}

type LaptopsCacheRepo interface {
	Set(ctx context.Context, key string, value []*models.Laptop) error
	Get(ctx context.Context, key string) ([]*models.Laptop, error)
}

