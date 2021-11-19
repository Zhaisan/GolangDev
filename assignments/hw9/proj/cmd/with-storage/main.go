package main

import (
	"context"
	"lectures-6/internal/cache/redis_cache"
	"lectures-6/internal/http"
	"lectures-6/internal/store/postgres"
	"log"
)

const (
	cacheDB = 1
	cacheExpTime = 1500
	cachePort = "localhost:6379"
)

func main() {
	urlExample := "postgres://localhost:5431/laptopsbd"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	//cache, err := lru.New2Q(6)
	//if err != nil {
	//	panic(err)
	//}

	cache := redis_cache.NewRedisCache(cachePort, cacheDB, cacheExpTime)

	srv := http.NewServer(context.Background(),
		http.WithAddress(":8075"),
		http.WithStore(store),
		http.WithCache(cache),
	)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
}
