package factory

import (
	"github.com/taisa831/go-ddd/domain/factory"
	"github.com/taisa831/go-ddd/domain/model"
)

type CircleFactory struct{}

func NewCircleFactory() factory.CircleFactory {
	return &CircleFactory{}
}

func (f *CircleFactory) Create(name *model.CircleName, owner *model.User) (*model.Circle, error) {
	conf := model.CircleConfig{
		Name:  name,
		Owner: owner,
	}
	circle, err := model.NewCircle(conf)
	if err != nil {
		return nil, err
	}
	return circle, nil
}
