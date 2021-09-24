package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taisa831/go-ddd/infrastracture/repository"
	"github.com/taisa831/go-ddd/infrastracture/service"
	"github.com/taisa831/go-ddd/interfaces/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
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

	router := gin.Default()

	r := repository.NewRepository(db)
	us := service.NewUserService(r)

	uh := handler.NewUserHandler(r, us)
	router.GET("/users", uh.List)
	router.POST("/users", uh.Create)
	router.Run()
}
