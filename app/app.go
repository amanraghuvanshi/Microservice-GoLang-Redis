package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

type App struct {
	router http.Handler
	// The redis client
	rdb *redis.Client
}

// constructor for the application

func New() *App {
	// initializing the client
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping().Err()
	if err != nil {
		return fmt.Errorf("redis connection failed, %w", err)
	}

	fmt.Println("Starting the server!!")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server!", err)
	}
	return nil
}
