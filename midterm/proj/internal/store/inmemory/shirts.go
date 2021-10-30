package inmemory

import (
	"context"
	"fmt"
	"lectures-6/internal/models"
	"sync"
)

type ShirtsRepo struct {
	data map[int]*models.Shirt

	mu *sync.RWMutex
}

func (db *ShirtsRepo) Create(ctx context.Context, shirt *models.Shirt) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[shirt.ID] = shirt
	return nil
}

func (db *ShirtsRepo) All(ctx context.Context) ([]*models.Shirt, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	shirts := make([]*models.Shirt, 0, len(db.data))
	for _, shirt := range db.data {
		shirts = append(shirts, shirt)
	}

	return shirts, nil
}

func (db *ShirtsRepo) ByID(ctx context.Context, id int) (*models.Shirt, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	shirt, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No shirt with id %d", id)
	}

	return shirt, nil
}

func (db *ShirtsRepo) Update(ctx context.Context, shirt *models.Shirt) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[shirt.ID] = shirt
	return nil
}

func (db *ShirtsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
