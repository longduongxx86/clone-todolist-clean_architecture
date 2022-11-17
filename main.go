package main

import (
	_ "first-app/docs"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	todotrpt "first-app/module/item/transport"
)

// @title         Todolist 
// @version       1.0
// @description   Todolist API using Swaggo
// @contact.name  DuongBaoLong
// @contact.email longduongxx86@gmail.com
// @host          localhost:8080
// @BasePath      /
func main() {
	dsn := "host=localhost user=postgres password=Nguyenthuyngan00 dbname=todolist port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to PostgreSQL:", err)
	}
 
	router := gin.Default()

	v1 := router.Group("")
	{
		
		v1.POST("/todos", todotrpt.HandleCreateItem(db))

		v1.GET("/todos", todotrpt.HandleListItem(db))

		v1.GET("/todos/:id", todotrpt.HandleFindAnItem(db))
	
		v1.PUT("/todos/:id", todotrpt.HandleUpdateAnItem(db))

		v1.DELETE("/todos/:id", todotrpt.HandleDeleteAnItem(db))

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
