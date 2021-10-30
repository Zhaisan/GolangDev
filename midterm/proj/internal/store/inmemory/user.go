package inmemory

import (
	"context"
	"errors"
	"fmt"
	"lectures-6/internal/models"
	"sync"
)

type UsersRepo struct {
	data map[string]*models.User

	mu *sync.RWMutex
}

func (db *UsersRepo) Create(ctx context.Context, user *models.User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if !user.IsEmailValid() {
		return errors.New("Invalid email error!")
	}
	
	db.data[user.Email] = user
	return nil
}

func (db *UsersRepo) All(ctx context.Context) ([]*models.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	users := make([]*models.User, 0, len(db.data))
	for _, user := range db.data {
		users = append(users, user)
	}

	return users, nil
}

func (db *UsersRepo) ByID(ctx context.Context, email string) (*models.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	user, ok := db.data[email]
	if !ok {
		return nil, fmt.Errorf("No user with id %v", email)
	}

	return user, nil
}

