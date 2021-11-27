package repository

import "github.com/taisa831/go-ddd/domain/model"

type Repository interface {
	FindUserByName(string) (*model.User, error)
	FindUserByID(string) (*model.User, error)
	FindUsers() ([]*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUser(*model.User) error

	SaveCircle(*model.Circle) error
	FindCircleByID(*model.CircleID) (*model.Circle, error)
	FindCircleByName(*model.CircleName) (*model.Circle, error)
}
