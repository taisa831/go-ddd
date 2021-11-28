package usecase

import (
	"errors"

	"github.com/taisa831/go-ddd/domain/factory"
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
)

type CircleUsecase struct {
	f factory.CircleFactory
	s service.CircleService
	r repository.Repository
}

func NewCircleUsecase(f factory.CircleFactory, s service.CircleService, rep repository.Repository) *CircleUsecase {
	return &CircleUsecase{
		f, s, rep,
	}
}

func (u *CircleUsecase) Create(userID, circleName string) (*model.Circle, error) {
	user, err := u.r.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	cName, err := model.NewCircleName(circleName)
	if err != nil {
		return nil, err
	}

	circle, err := u.f.Create(cName, user)
	if err != nil {
		return nil, err
	}

	exist, err := u.s.Exists(circle)
	if err != nil {
		return nil, err
	}

	if exist {
		return nil, errors.New("duplicated")
	}

	if err := u.r.SaveCircle(circle); err != nil {
		return nil, err
	}

	return circle, nil
}

func (u *CircleUsecase) Join(userID, circleID string) error {
	user, err := u.r.FindUserByID(userID)
	if err != nil {
		return err
	}

	cID, err := model.NewCircleID(circleID)
	if err != nil {
		return err
	}

	circle, err := u.r.FindCircleByID(cID)
	if err != nil {
		return err
	}

	if circle == nil {
		return errors.New("circle is empty")
	}

	// if len(circle.Members) >= 29 {
	if circle.IsFull() {
		return errors.New("circle is full")
	}

	// circle.Members = append(circle.Members, user)
	circle.Join(user)

	return u.r.SaveCircle(circle)
}
