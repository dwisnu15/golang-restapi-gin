package main

import (
	controller "GinAPI/controllers"
	"GinAPI/repositories"
	service "GinAPI/services"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func inject(db *sql.DB) (*gin.Engine, error) {
	itemRepository := repositories.InitItemRepo(db)
	itemService := service.CreateItemsService(itemRepository)

	itemController, err := controller.InitItemsController(itemService)
	if err != nil {
		return nil, err
	}
	router, err := itemController.SetupRouter()
	if err != nil {
		return nil, err
	}
	return router, nil

}
