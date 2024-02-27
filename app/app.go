package app

import (
	"context"
	"log"
	"net/http"
)

type App struct {
	router http.Handler
}

// constructor for the application

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server!", err)
	}
	return nil
}
