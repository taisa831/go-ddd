package model

import (
	"errors"

	"github.com/google/uuid"
)

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
	Owner   string
	Members []*CircleMember
}

type CircleConfig struct {
	Name    *CircleName
	Owner   *User
	Members []*CircleMember
}

func NewCircle(conf CircleConfig) (*Circle, error) {
	if conf.Name == nil {
		return nil, errors.New("")
	}

	uuid, err := NewCircleID(uuid.NewString())
	if err != nil {
		return nil, err
	}

	circle := &Circle{
		ID:      *uuid,
		Name:    *conf.Name,
		Owner:   conf.Owner.ID,
		Members: conf.Members,
	}

	circleMember := NewCircleMember(conf.Owner, circle.ID)
	circle.Members = append(circle.Members, circleMember)

	return circle, nil
}

func (c *Circle) IsFull() bool {
	return c.countMembers() >= 30
}

func (c *Circle) countMembers() int {
	return len(c.Members) + 1
}

// デメテルの法則
func (c *Circle) Join(user *User) {
	c.Members = append(c.Members, NewCircleMember(user, c.ID))
}
