package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"time"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(app.RateLimitMiddleware)

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsOptions).Handler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/users", app.handleRegisterUser)
		r.Post("/users/activate", app.handleActivateUser)
		r.Post("/token/auth", app.handleLoginUser)
		r.Post("/token/refresh", app.handleRefreshToken)
	})

	r.Route("/api/v1/contacts", func(r chi.Router) {
		r.Use(app.AuthMiddleware)
		r.Get("/", app.handleGetAllContact)
		r.Post("/", app.handleCreateContact)
		r.Get("/{id}", app.handleGetContact)
		r.Patch("/{id}", app.handleUpdateContact)
		r.Delete("/{id}", app.handleDeleteContact)
	})

	return r
}
