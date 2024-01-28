package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // X-CSRF-Token is added to avoid CSRF attacks
		ExposedHeaders:   []string{"Link"},                                                    // Only allow headers that we're ok with being public
		AllowCredentials: true,                                                                // Allows sending cookies
		MaxAge:           300,                                                                 // Maximum value not ignored by any of major browsers
	}))

	mux.Use(middleware.Heartbeat(("/ping")))

	mux.Post("/", app.Broker)

	return mux
}
