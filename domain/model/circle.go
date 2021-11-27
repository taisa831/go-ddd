package model

import "errors"

type CircleID struct {
	Value string
}

func NewCircleID(val string) (*CircleID, error) {
	if val == "" {
		return nil, errors.New("")
	}
	return &CircleID{
		Value: val,
	}, nil
}

type CircleName struct {
	Value string
}

func NewCircleName(val string) (*CircleName, error) {
	if val == "" {
		return nil, errors.New("")
	}

	if len(val) < 3 || len(val) > 20 {
		return nil, errors.New("")
	}

	return &CircleName{
		Value: val,
	}, nil
}

type Circle struct {
	ID      CircleID
	Name    CircleName
	Owner   User
	Members []*User
}

type CircleConfig struct {
	ID      *CircleID
	Name    *CircleName
	Owner   *User
	Members []*User
}

func NewCircle(conf CircleConfig) (*Circle, error) {
	if conf.ID == nil || conf.Name == nil {
		return nil, errors.New("")
	}

	return &Circle{
		ID:      *conf.ID,
		Name:    *conf.Name,
		Owner:   *conf.Owner,
		Members: conf.Members,
	}, nil
}
