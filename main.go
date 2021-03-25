package main

import (
	"GinAPI/configs"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//print current date and time
	//os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())
	//i should create a function which will embed all of...
	db :=configs.InitDBConnection()

	router, err := InjectRouter(db)
	if err != nil {
		log.Fatalf("Failed on router settings: %v\n", err)
	}

	//...this into my router
	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	//Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//this blocks until a signal is passed into the quit channel
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
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
