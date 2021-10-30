package inmemory

import (
	"context"
	"fmt"
	"lectures-6/internal/models"
	"sync"
)

type ToysRepo struct {
	data map[int]*models.Toy

	mu *sync.RWMutex
}

func (db *ToysRepo) Create(ctx context.Context, toy *models.Toy) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[toy.ID] = toy
	return nil
}

func (db *ToysRepo) All(ctx context.Context) ([]*models.Toy, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	toys := make([]*models.Toy, 0, len(db.data))
	for _, toy := range db.data {
		toys = append(toys, toy)
	}

	return toys, nil
}

func (db *ToysRepo) ByID(ctx context.Context, id int) (*models.Toy, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	toy, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No toy with id %d", id)
	}

	return toy, nil
}

func (db *ToysRepo) Update(ctx context.Context, toy *models.Toy) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[toy.ID] = toy
	return nil
}

func (db *ToysRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
