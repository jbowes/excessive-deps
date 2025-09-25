package handlers

import (
	"encoding/json"
	"net/http"

	"sampleapp/internal/models"
)

func (h *Handlers) ListTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.todos.List())
}

func (h *Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid payload")
		return
	}
	created, err := h.todos.Create(t)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
