package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Info struct {
	ver  string
	lang string
}

type Server struct {
	engine *gin.Engine
	info   Info
}

func New() Server {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return Server{
		engine: r,
		info: Info{
			ver:  "0.0.1",
			lang: "en",
		},
	}
}

func (s *Server) Info() string {
	return s.info.ver
}

func (s *Server) SetVer(ver string) {
	s.info.ver = ver
}

func (s *Server) Run() error {
	return s.engine.Run()
}
