package main

import (
	"GinAPI/configs"
	repository "GinAPI/repositories"
	controller "GinAPI/controllers"
	service "GinAPI/services"
	"GinAPI/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	//print current date and time
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())

	db :=configs.InitDBConnection()
	itemsRepo := repository.InitItemRepo(db)
	itemsService := service.CreateItemsService(itemsRepo)

	router := gin.Default()
	router.Use(middlewares.Logger())
	controller.InitItemsController(router, itemsService)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

////initPostgres
//models.InitGormPostgres()
//defer models.DB.Close()
//
////set router with gin
//gin.SetMode(gin.ReleaseMode)
//router := gin.Default()
//
////setup route group for the api
//api := router.Group("/api")
//
////find all
//api.GET("/items", controllers.FindAllItems)
////find by id
//api.GET("/items/:id", controllers.FindItemById)
////create new item
//api.POST("/items", controllers.AddItemsInput)
////update item
//api.PATCH("/items/:id", controllers.UpdateItem)
////delete item
//api.DELETE("/items/:id", controllers.DeleteItem)
////start server on port 8080
//router.Run(":8080")
