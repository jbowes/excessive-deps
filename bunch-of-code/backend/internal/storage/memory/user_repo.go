package memory

import "sampleapp/internal/models"

type UserRepo struct {
	items []models.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{items: make([]models.User, 0, 16)}
}

func (r *UserRepo) List() []models.User {
	out := make([]models.User, len(r.items))
	copy(out, r.items)
	return out
}

func (r *UserRepo) Create(u models.User) (models.User, error) {
	r.items = append(r.items, u)
	return u, nil
}
