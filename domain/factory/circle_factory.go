package factory

import "github.com/taisa831/go-ddd/domain/model"

type CircleFactory interface {
	Create(name *model.CircleName, owner *model.User) (*model.Circle, error)
}
