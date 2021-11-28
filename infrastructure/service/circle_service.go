package service

import (
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
	"gorm.io/gorm"
)

type CircleService struct {
	r repository.Repository
}

func NewCircleService(rep repository.Repository) service.CircleService {
	return &CircleService{
		r: rep,
	}
}

func (s *CircleService) Exists(circle *model.Circle) (bool, error) {
	dup, err := s.r.FindCircleByName(&circle.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return dup != nil, nil
}
