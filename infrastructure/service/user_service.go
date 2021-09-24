package service

import (
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
	_, err := s.r.FindUserByName(name)
	if err != nil {
		return false, err
	}
	return true, nil
}
