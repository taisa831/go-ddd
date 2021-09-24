package service

type UserService interface {
	Exists(name string) (bool, error)
}
