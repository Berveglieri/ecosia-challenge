package router

import (
	"ecosia/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/status", handlers.StatusHandler)
	router.Get("/tree", handlers.TreeHandler)
	router.Get("/tree/{name}", handlers.TreeHandlerWithParams)

	return router
}
