package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Vafelence/fifa_project/internal/handlers"
	"github.com/Vafelence/fifa_project/internal/pkg/server"
	"github.com/Vafelence/fifa_project/internal/pkg/sql"
	"github.com/Vafelence/fifa_project/internal/services"
	"github.com/Vafelence/fifa_project/internal/storage"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "app", "fifa-service")

	db, err := sql.New()
	if err != nil {
		fmt.Println("can't start db", err)
		os.Exit(1)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("can't close db", err)
		}
	}()
	dataStorage := storage.New(db)
	srv := server.New()
	router := srv.Router()
	service := services.New(dataStorage)
	handler := handlers.New(service)

	router.POST("/player/*id", handler.Player)
	router.POST("/player/*id/delete", handler.Delete)

	router.GET("/hello", handler.Hello)

	err = srv.Run(ctx)
	if err != nil {
		fmt.Println("can't start server", err)
		os.Exit(1)
	}
}
