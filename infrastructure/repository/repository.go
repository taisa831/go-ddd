package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbRepository struct {
	db *gorm.DB
}

func OpenDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("SCHEMA"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(151)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewRepository(db *gorm.DB) repository.Repository {
	return &dbRepository{
		db: db,
	}
}
