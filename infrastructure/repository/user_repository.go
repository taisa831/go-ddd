package repository

import (
	"github.com/taisa831/go-ddd/domain/model"
)

func (r *dbRepository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *dbRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *dbRepository) FindUsers() ([]*model.User, error) {
	users := []*model.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}