package repository

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/taisa831/go-ddd/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var rdb *gorm.DB

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env ファイルの読み込みに失敗しました。")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("SCHEMA"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(151)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	rdb = db
	code := m.Run()
	os.Exit(code)
}

func truncate(db *gorm.DB) {
	db.Exec("truncate users")
}
