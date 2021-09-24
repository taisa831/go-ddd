package service

import (
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
)

type UserService struct {
	r repository.Repository
}

func NewUserService(r repository.Repository) service.UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) Exists(name string) (bool, error) {
	user, err := s.r.FindUserByName(name)
	if err != nil && err != model.ErrNotFound {
		return false, err
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}
