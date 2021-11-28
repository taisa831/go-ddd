package model

import "github.com/google/uuid"

type CircleMember struct {
	ID       string
	CircleID CircleID
	UserID   string
}

func NewCircleMember(u *User, circleID CircleID) *CircleMember {
	return &CircleMember{
		ID:       uuid.NewString(),
		CircleID: circleID,
		UserID:   u.ID,
	}
}
