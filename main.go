package main

import (
	"GinAPI/controllers"
	"GinAPI/middlewares"
	"GinAPI/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	//print current date and time
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())

	models.InitDBConnection()
	defer models.DB.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//var log = logrus.New()
	//router.Use(middlewares.Logger(log))

	api := router.Group("/api")
	api.Use(middlewares.Logger())
	api.GET("/items", controllers.ListAllItems)//tested
	api.GET("/items/:id", controllers.FindItemById)//tested
	api.POST("/items", controllers.InsertItem)//tested
	api.DELETE("/items/:id", controllers.DeleteItem)//tested
	api.PATCH("/items/:id", controllers.PatchItem)//tested
	router.Run(":8080")

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
