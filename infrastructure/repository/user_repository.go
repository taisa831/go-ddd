package repository

import (
	"github.com/taisa831/go-ddd/domain/model"
)

func (r *rdbRepository) FindUserByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *rdbRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}
