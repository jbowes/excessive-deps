package handlers

import (
	"encoding/json"
	"net/http"

	"sampleapp/internal/services"
)

type Handlers struct {
	users services.UserService
	todos services.TodoService
}

func NewHandlers(users services.UserService, todos services.TodoService) *Handlers {
	return &Handlers{users: users, todos: todos}
}

func (h *Handlers) writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
