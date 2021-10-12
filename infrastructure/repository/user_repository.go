package repository

import (
	"github.com/taisa831/go-ddd/domain/model"
)

type User struct {
	ID      string
	Name    string
	Address string
}

func (r *dbRepository) FindUserByName(name string) (*model.User, error) {
	var user User
	if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return r.convertToUserModel(&user), nil
}

func (r *dbRepository) FindUserByID(id string) (*model.User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return r.convertToUserModel(&user), nil
}

func (r *dbRepository) CreateUser(user *model.User) error {
	return r.db.Create(r.convertToUserRecord(user)).Error
}

func (r *dbRepository) FindUsers() ([]*model.User, error) {
	users := []*User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return r.convertToUserModels(users), nil
}

func (r *dbRepository) UpdateUser(user *model.User) error {
	return r.db.Model(User{}).Where("id = ?", user.ID).Updates(User{
		Name:    user.Name,
		Address: user.Address,
	}).Error
}

func (r *dbRepository) DeleteUser(user *model.User) error {
	return r.db.Model(User{}).Where("id = ?", user.ID).Delete(user).Error
}

func (r *dbRepository) convertToUserRecord(user *model.User) *User {
	return &User{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
	}
}

func (r dbRepository) convertToUserModel(user *User) *model.User {
	return &model.User{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
	}
}

func (r dbRepository) convertToUserModels(users []*User) []*model.User {
	userModels := make([]*model.User, len(users))
	for i, u := range users {
		userModels[i] = &model.User{
			ID:      u.ID,
			Name:    u.Name,
			Address: u.Address,
		}
	}
	return userModels
}
