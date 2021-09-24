package model

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

type UserCreateConfig struct {
	Name string
}

func NewUser(conf UserCreateConfig) (*User, error) {
	u := &User{
		ID: uuid.NewString(),
	}
	err := u.ChangeName(conf.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
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

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName, lastName string) (*FullName, error) {
	if firstName == "" {
		return nil, errors.New("firstName required")
	}

	if lastName == "" {
		return nil, errors.New("firstName required")
	}

	return &FullName{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

func (m *FullName) FirstName() string {
	return m.firstName
}

func (m *FullName) LastName() string {
	return m.lastName
}

func PrintFullName() error {
	fullName, err := NewFullName("taro", "suzuki")
	if err != nil {
		return err
	}
	fmt.Println(fullName)
	return nil
}
