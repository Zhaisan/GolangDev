package main

import (
	"context"
	"lectures-6/internal/http"
	"lectures-6/internal/store/postgres"
	"log"
)

func main() {
	urlExample := "postgres://localhost:5431/laptopsbd"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	srv := http.NewServer(context.Background(), ":8075", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
}
