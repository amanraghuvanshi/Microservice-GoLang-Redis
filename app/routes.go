package app

import (
	"net/http"

	"github.com/amanraghuvanshi/microservice-golang-redis/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// setting up the subrouter that will automatically add the orders path to the url

	router.Route("/orders", loadOrderRouters)

	return router
}

func loadOrderRouters(router chi.Router) {
	orderHandler := &handlers.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/update/{id}", orderHandler.UpdateByID)
	router.Delete("/delete/{id}", orderHandler.DeleteByID)
}
