package repository

import (
	"github.com/taisa831/go-ddd/domain/repository"
	"gorm.io/gorm"
)

type rdbRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Repository {
	return &rdbRepository{
		db: db,
	}
}
