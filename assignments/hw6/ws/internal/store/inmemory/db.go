package inmemory

import (
	"context"
	"fmt"
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"sync"
)

type DB struct {
	data map[int]*models.Item

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		data: make(map[int]*models.Item),
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Create(ctx context.Context, laptop *models.Item) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[laptop.ID] = laptop
	return nil
}

func (db *DB) All(ctx context.Context) ([]*models.Item, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	items := make([]*models.Item, 0, len(db.data))
	for _, item := range db.data {
		items = append(items, item)
	}

	return items, nil
}

func (db *DB) ByID(ctx context.Context, id int) (*models.Item, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	item, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No item with id %d", id)
	}

	return item, nil
}

func (db *DB) Update(ctx context.Context, item *models.Item) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[item.ID] = item
	return nil
}

func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
