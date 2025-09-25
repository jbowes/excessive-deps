package services

import (
	"errors"
	"strings"
	"time"

	"sampleapp/internal/models"
	"sampleapp/internal/storage/memory"
	"sampleapp/pkg/utils"
)

type UserService interface {
	List() []models.User
	Create(models.User) (models.User, error)
}

type userService struct {
	repo *memory.UserRepo
}

func NewUserService() UserService {
	return &userService{repo: memory.NewUserRepo()}
}

func (s *userService) List() []models.User {
	return s.repo.List()
}

func (s *userService) Create(in models.User) (models.User, error) {
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))
	in.Name = strings.TrimSpace(in.Name)
	if !utils.IsValidEmail(in.Email) {
		return models.User{}, errors.New("invalid email")
	}
	if in.Name == "" {
		return models.User{}, errors.New("name required")
	}
	in.ID = utils.NewID()
	in.CreatedAt = time.Now()
	return s.repo.Create(in)
}
