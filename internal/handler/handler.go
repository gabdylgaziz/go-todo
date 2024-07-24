package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gotodo/internal/handler/http"
	"gotodo/internal/service/gotodo"
	"gotodo/pkg/server/router"
)

type Dependencies struct {
	TodoService *gotodo.Service
}

type Configuration func(h *Handler) error

type Handler struct {
	dependencies Dependencies

	HTTP *chi.Mux
}

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	h = &Handler{
		dependencies: d,
	}

	for _, cfg := range configs {
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		h.HTTP = router.New()

		h.HTTP.Use(middleware.Timeout(60))
		h.HTTP.Get("/swagger/*", httpSwagger.WrapHandler)

		todoHandler := http.NewTodoHandler(h.dependencies.TodoService)

		h.HTTP.Get("/health", todoHandler.HealthCheck)

		h.HTTP.Route("/api/todo-list", func(r chi.Router) {
			r.Mount("/tasks", todoHandler.Routes())
		})

		return
	}
}
