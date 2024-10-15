package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func rateLimit(next http.Handler) http.Handler {

	limiter := rate.NewLimiter(10, 1)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(rateLimit)

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsOptions).Handler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/users", app.handleRegisterUser)
		r.Post("/users/activate", app.handleActivateUser)
		r.Post("/token/auth", app.handleLoginUser)
	})

	r.Route("/api/v1/contacts", func(r chi.Router) {
		r.Use(app.AuthMiddleware)
		r.Post("/", app.handleCreateContact)
		r.Get("/{id}", app.handleGetContact)
	})

	return r
}
