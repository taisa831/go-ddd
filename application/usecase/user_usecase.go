package usecase

import (
	"fmt"

	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
	"github.com/taisa831/go-ddd/interfaces/request"
)

type UserUsecase struct {
	r  repository.Repository
	us service.UserService
}

func NewUserUsecase(r repository.Repository, us service.UserService) *UserUsecase {
	return &UserUsecase{
		r:  r,
		us: us,
	}
}
func (u *UserUsecase) Primitive() {
	fullName := "taro suzuki"
	fmt.Println(fullName)
}

func (u *UserUsecase) Create(req request.UserCreateRequest) error {
	b, err := u.us.Exists(req.Name)
	if err != nil {
		return err
	}
	if b {
		return &model.UserExistsError{}
	}

	conf := model.UserCreateConfig{
		Name: req.Name,
	}
	user, err := model.NewUser(conf)
	if err != nil {
		return err
	}
	if err := u.r.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) List() ([]*model.User, error) {
	return u.r.FindUsers()
}

func (u *UserUsecase) Update(userID string, req request.UserUpdateRequest) error {
	duplicated, err := u.us.Exists(req.Name)
	if err != nil {
		return err
	}

	if duplicated {
		return &model.UserExistsError{}
	}

	user, err := u.r.FindUserByID(userID)
	if err != nil {
		return err
	}

	conf := model.UserUpdateConfig{
		Name:    req.Name,
		Address: req.Address,
	}

	if err := user.Update(conf); err != nil {
		return err
	}

	if err := u.r.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Get(userID string) (*model.User, error) {
	return u.r.FindUserByID(userID)
}

func (u *UserUsecase) Delete(userID string) error {
	user, err := u.r.FindUserByID(userID)
	if err != nil {
		return err
	}

	if err := u.r.DeleteUser(user); err != nil {
		return err
	}
	return nil
}
