package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func (s *Server) Router() gin.IRouter {
	return s.engine
}

func (s *Server) Info() string {
	return s.info.ver
}

func (s *Server) SetVer(ver string) {
	s.info.ver = ver
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "can't shutdown server")
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Printf("Server %s exiting", ctx.Value("app").(string))
	return nil
}
