package repository

import "github.com/taisa831/go-ddd/domain/model"

type Repository interface {
	FindUserByName(name string) (*model.User, error)
	FindUsers() ([]*model.User, error)
	CreateUser(user *model.User) error
}
