package services

import (
	"errors"
	"strings"
	"time"

	"sampleapp/internal/models"
	"sampleapp/internal/storage/memory"
	"sampleapp/pkg/utils"
)

type TodoService interface {
	List() []models.Todo
	Create(models.Todo) (models.Todo, error)
}

type todoService struct {
	repo *memory.TodoRepo
}

func NewTodoService() TodoService {
	return &todoService{repo: memory.NewTodoRepo()}
}

func (s *todoService) List() []models.Todo {
	return s.repo.List()
}

func (s *todoService) Create(in models.Todo) (models.Todo, error) {
	in.Title = strings.TrimSpace(in.Title)
	if in.Title == "" {
		return models.Todo{}, errors.New("title required")
	}
	if in.UserID == "" {
		return models.Todo{}, errors.New("userId required")
	}
	in.ID = utils.NewID()
	in.CreatedAt = time.Now()
	return s.repo.Create(in)
}
