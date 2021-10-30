package inmemory

import (
	"context"
	"fmt"
	"lectures-6/internal/models"
	"sync"
)

type SnowboardsRepo struct {
	data map[int]*models.Snowboard

	mu *sync.RWMutex
}

func (db *SnowboardsRepo) Create(ctx context.Context, snowboard *models.Snowboard) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[snowboard.ID] = snowboard
	return nil
}

func (db *SnowboardsRepo) All(ctx context.Context) ([]*models.Snowboard, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	snowboards := make([]*models.Snowboard, 0, len(db.data))
	for _, snowboard := range db.data {
		snowboards = append(snowboards, snowboard)
	}

	return snowboards, nil
}

func (db *SnowboardsRepo) ByID(ctx context.Context, id int) (*models.Snowboard, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	snowboard, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No snowboard with id %d", id)
	}

	return snowboard, nil
}

func (db *SnowboardsRepo) Update(ctx context.Context, snowboard *models.Snowboard) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[snowboard.ID] = snowboard
	return nil
}

func (db *SnowboardsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
