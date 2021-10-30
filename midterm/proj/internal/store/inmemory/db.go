package inmemory

import (
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"sync"
)

type DB struct {
	laptopsRepo store.LaptopsRepository
	snowboardsRepo store.SnowboardsRepository
	shirtsRepo store.ShirtsRepository
	toysRepo store.ToysRepository
	usersRepo store.UsersRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}

func (db *DB) Laptops() store.LaptopsRepository {
	if db.laptopsRepo == nil {
		db.laptopsRepo = &LaptopsRepo{
			data: make(map[int]*models.Laptop),
			mu:   new(sync.RWMutex),
		}
	}

	return db.laptopsRepo
}

func (db *DB) Snowboards() store.SnowboardsRepository {
	if db.snowboardsRepo == nil {
		db.snowboardsRepo = &SnowboardsRepo{
			data: make(map[int]*models.Snowboard),
			mu:   new(sync.RWMutex),
		}
	}

	return db.snowboardsRepo
}

func (db *DB) Shirts() store.ShirtsRepository {
	if db.shirtsRepo == nil {
		db.shirtsRepo = &ShirtsRepo{
			data: make(map[int]*models.Shirt),
			mu:   new(sync.RWMutex),
		}
	}

	return db.shirtsRepo
}

func (db *DB) Toys() store.ToysRepository {
	if db.toysRepo == nil {
		db.toysRepo = &ToysRepo{
			data: make(map[int]*models.Toy),
			mu:   new(sync.RWMutex),
		}
	}

	return db.toysRepo
}

func (db *DB) Users() store.UsersRepository {
	if db.usersRepo == nil {
		db.usersRepo = &UsersRepo{
			data: make(map[string]*models.User),
			mu:   new(sync.RWMutex),
		}
	}

	return db.usersRepo
}