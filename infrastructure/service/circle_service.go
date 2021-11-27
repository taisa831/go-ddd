package service

import (
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
)

type CircleService struct {
	rep repository.Repository
}

func NewCircleService(rep repository.Repository) service.CircleService {
	return &CircleService{
		rep: rep,
	}
}

func (s *CircleService) Exists(circle *model.Circle) (bool, error) {
	dup, err := s.rep.FindCircleByName(&circle.Name)
	if err != nil {
		return false, err
	}
	return dup != nil, nil
}
