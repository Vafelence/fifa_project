package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Vafelence/fifa_project/internal/pkg/server"
	"github.com/Vafelence/fifa_project/internal/pkg/sql"
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

	router.POST("/player/*id", func(c *gin.Context) {
		ctx := c.Request.Context()
		playerID := strings.ReplaceAll(c.Param("id"), "/", "")

		type Player struct {
			Name string `json:"name" binding:"required"`
		}
		var player Player
		err := c.BindJSON(&player)
		if err != nil {
			fmt.Printf("can't bind player data %s: %v", err, playerID)
			c.JSON(http.StatusInternalServerError, &gin.H{
				"status": "error",
				"err":    err,
			})
			return
		}

		if playerID == "" {
			err = dataStorage.AddPlayer(ctx, player.Name)
			if err != nil {
				fmt.Printf("can't add to storage player data %s: %v", err, player)
				c.JSON(http.StatusInternalServerError, &gin.H{
					"status": "error",
					"err":    err,
				})
				return
			}
		} else {
			err = dataStorage.UpdatePlayer(ctx, playerID, player.Name)
			if err != nil {
				fmt.Printf("can't update to storage player data %s: %v", err, player)
				c.JSON(http.StatusInternalServerError, &gin.H{
					"status": "error",
					"err":    err,
				})
				return
			}
		}

		c.JSON(http.StatusOK, &gin.H{
			"status": "ok",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", &gin.H{
			"page_title":    "Выбор игровых команд",
			"random_team_1": "Одутловатые",
			"random_team_2": "Отрафиты",
		})
	})

	err = srv.Run(ctx)
	if err != nil {
		fmt.Println("can't start server", err)
		os.Exit(1)
	}
}
