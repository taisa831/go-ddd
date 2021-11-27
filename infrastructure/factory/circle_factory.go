package factory

import "github.com/taisa831/go-ddd/domain/model"

type CircleFactory struct{}

func (f *CircleFactory) Create(name *model.CircleName, owner *model.User) (*model.Circle, error) {
	return nil, nil
}
