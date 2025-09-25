package handlers

import (
	"encoding/json"
	"net/http"

	"sampleapp/internal/models"
)

func (h *Handlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.users.List())
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid payload")
		return
	}
	created, err := h.users.Create(u)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
