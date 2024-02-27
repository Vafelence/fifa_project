package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Vafelence/fifa_project/internal/pkg/server"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "app", "fifa-service")
	srv := server.New()

	err := srv.Run(ctx)
	if err != nil {
		fmt.Println("can't start server", err)
		os.Exit(1)
	}
}
