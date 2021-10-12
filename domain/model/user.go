package model

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID      string
	Name    string
	Address string
}

type UserCreateConfig struct {
	Name    string
	Address string
}

type UserUpdateConfig struct {
	Name    string
	Address string
}

func NewUser(conf UserCreateConfig) (*User, error) {
	u := &User{
		ID:      uuid.NewString(),
		Address: conf.Address,
	}
	err := u.ChangeName(conf.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *User) Update(conf UserUpdateConfig) error {
	if conf.Name == "" {
		return fmt.Errorf("name is required")
	}

	if conf.Address == "" {
		return fmt.Errorf("address is required")
	}
	m.Name = conf.Name
	m.Address = conf.Address
	return nil
}

func (m *User) ChangeName(name string) error {
	if name == "" {
		return fmt.Errorf("ユーザ名は必須です。")
	}
	if len(name) < 3 {
		return fmt.Errorf("ユーザ名は3文字以上です。%s", name)
	}
	m.Name = name
	return nil
}
