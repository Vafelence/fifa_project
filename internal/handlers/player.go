package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Vafelence/fifa_project/internal/services"
)

func (h *Handler) Player(c *gin.Context) {
	ctx := c.Request.Context()
	ctx.Value("user-uuid")
	playerID := strings.ReplaceAll(c.Param("id"), "/", "")

	type player struct {
		Name string `json:"name" binding:"required"`
	}
	var pl player
	err := c.BindJSON(&pl)
	if err != nil {
		fmt.Printf("can't bind player data %s: %v", err, playerID)
		c.JSON(http.StatusInternalServerError, &gin.H{
			"status": "error",
			"err":    err,
		})
		return
	}

	err = h.service.PlayerManagement(ctx, services.Player{
		ID:   playerID,
		Name: pl.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"status": "error",
			"err":    err,
		})
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"status": "ok",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	playerID := strings.ReplaceAll(c.Param("id"), "/", "")

	h.service.Delete(ctx, playerID)
	return
}
