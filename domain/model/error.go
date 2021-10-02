package model

type UserExistsError struct{}

func (e *UserExistsError) Error() string {
	return "ユーザーは存在します。"
}
