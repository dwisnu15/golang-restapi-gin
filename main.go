package main

import (
	"GinAPI/configs"
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db :=configs.InitDBConnection()
	router, err := inject(db)
	if err != nil {
		log.Fatalf("Failed injecting router: %v\n", err)
	}

	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("%v\n", err)
	}
	// Retrieved from Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()
	log.Printf("Listening on port %v\n", srv.Addr)

	//create kill signal of channel
	stopServer := make(chan os.Signal, 1)

	signal.Notify(stopServer, syscall.SIGINT, syscall.SIGTERM)

	//blocks until a signal is sent to stop server channel
	<-stopServer

	//inform context that it has 7 seconds to finish
	//handling received request
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()


	//shut down database source
	if err := db.Close(); err != nil {
		log.Fatalf("Shutting down db on problem: %v\n", err)
	}

	//shutdown api (server)
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
