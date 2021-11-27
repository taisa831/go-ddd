package service

import "github.com/taisa831/go-ddd/domain/model"

type CircleService interface {
	Exists(*model.Circle) (bool, error)
}
