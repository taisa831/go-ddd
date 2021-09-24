package model

import "errors"

var (
	ErrNotFound = errors.New("404")
)

type UserExistsError struct{}

func (e *UserExistsError) Error() string {
	return "ユーザーは存在します。"
}
