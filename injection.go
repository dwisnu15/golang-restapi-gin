package main

import (
	controller "GinAPI/controllers"
	"GinAPI/repositories"
	service "GinAPI/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

func InjectRouter(db *sql.DB) (*gin.Engine, error) {
	/*
	Repository
	 */
	itemRepository := repositories.InitItemRepo(db)

	/*
	Service
	 */
	itemsService := service.CreateItemsService(itemRepository)

	//init gin engine
	router := gin.Default()
	baseurl := viper.GetString("server.base_url")//api/
	timeouts := viper.GetInt("database.connection_timeout")
	maxSize := viper.GetInt64("database.max_upload_size")

	//router.Use(middlewares.Logger())

	controller.InitItemsController(&controller.ItemsContrConfig{
		R :              router,
		ItemsService:    itemsService,
		BaseURL:         baseurl,
		TimeoutDuration: time.Duration(timeouts) * time.Second, //set 5 seconds by default
		MaxBodyBytes:    maxSize,
	})

	return router, nil

}
