package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"sampleapp/internal/config"
	ih "sampleapp/internal/http/handlers"
	imw "sampleapp/internal/http/middleware"
	"sampleapp/internal/services"
	"sampleapp/pkg/logger"
)

func NewRouter(cfg config.Config, logg logger.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(imw.Logging(logg))
	r.Use(middleware.Recoverer)

	userSvc := services.NewUserService()
	todoSvc := services.NewTodoService()

	h := ih.NewHandlers(userSvc, todoSvc)

	r.Get("/health", h.Health)

	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(imw.Auth())
			r.Get("/users", h.ListUsers)
			r.Post("/users", h.CreateUser)
			r.Get("/todos", h.ListTodos)
			r.Post("/todos", h.CreateTodo)
		})
	})

	return r
}
