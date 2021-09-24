package repository

import (
	"github.com/taisa831/go-ddd/domain/model"
	"gorm.io/gorm"
)

func (r *dbRepository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *dbRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}
