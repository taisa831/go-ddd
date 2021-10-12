package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/taisa831/go-ddd/infrastructure/repository"
	"github.com/taisa831/go-ddd/infrastructure/service"
	"github.com/taisa831/go-ddd/interfaces/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env ファイルの読み込みに失敗しました。")
	}

	db, err := repository.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	router := gin.Default()
	r := repository.NewRepository(db)
	us := service.NewUserService(r)

	uh := handler.NewUserHandler(r, us)
	router.POST("/users", uh.Create)
	router.GET("/users", uh.List)
	router.GET("/users/:userId", uh.Get)
	router.PATCH("/users/:userId", uh.Update)
	router.DELETE("/users/:userId", uh.Delete)

	router.Run()
}
