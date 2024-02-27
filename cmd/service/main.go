package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Vafelence/fifa_project/internal/pkg/server"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "app", "fifa-service")
	srv := server.New()
	router := srv.Router()

	router.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", &gin.H{
			"page_title":    "Выбор игровых команд",
			"random_team_1": "Одутловатые",
			"random_team_2": "Отрафиты",
		})
	})

	err := srv.Run(ctx)
	if err != nil {
		fmt.Println("can't start server", err)
		os.Exit(1)
	}
}
