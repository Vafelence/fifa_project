package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Hello(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.service.Hello(ctx)
	c.HTML(http.StatusOK, "main.tmpl", data)
}
