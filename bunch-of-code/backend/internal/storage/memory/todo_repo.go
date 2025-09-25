package memory

import "sampleapp/internal/models"

type TodoRepo struct {
	items []models.Todo
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{items: make([]models.Todo, 0, 32)}
}

func (r *TodoRepo) List() []models.Todo {
	out := make([]models.Todo, len(r.items))
	copy(out, r.items)
	return out
}

func (r *TodoRepo) Create(t models.Todo) (models.Todo, error) {
	r.items = append(r.items, t)
	return t, nil
}
