package repository

import (
	"github.com/taisa831/go-ddd/domain/model"
)

func (r *dbRepository) FindUserByName(name string) (*model.User, error) {
	var user user
	if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return r.convertToUserModel(&user), nil
}

func (r *dbRepository) FindUserByID(id string) (*model.User, error) {
	var user user
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
	users := []*user{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return r.convertToUserModels(users), nil
}

func (r *dbRepository) UpdateUser(u *model.User) error {
	return r.db.Model(user{}).Where("id = ?", u.ID).Updates(user{
		Name:    u.Name,
		Address: u.Address,
	}).Error
}

func (r *dbRepository) DeleteUser(u *model.User) error {
	return r.db.Model(user{}).Where("id = ?", u.ID).Delete(u).Error
}

func (r *dbRepository) convertToUserRecord(u *model.User) *user {
	return &user{
		ID:      u.ID,
		Name:    u.Name,
		Address: u.Address,
	}
}

func (r dbRepository) convertToUserModel(u *user) *model.User {
	return &model.User{
		ID:      u.ID,
		Name:    u.Name,
		Address: u.Address,
	}
}

func (r dbRepository) convertToUserModels(us []*user) []*model.User {
	userModels := make([]*model.User, len(us))
	for i, u := range us {
		userModels[i] = &model.User{
			ID:      u.ID,
			Name:    u.Name,
			Address: u.Address,
		}
	}
	return userModels
}
