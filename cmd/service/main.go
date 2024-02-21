package main

import (
	"fmt"
	"os"

	"github.com/Vafelence/fifa_project/internal/pkg/server"
)

func main() {
	srv := server.New()
	fmt.Println(srv.Info())
	err := srv.Run()
	if err != nil {
		fmt.Println("can't start server", err)
		os.Exit(1)
	}
}
