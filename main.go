package main

import (
	"fmt"
	"os"
	"vault/internal/server"
)

func main() {
	server := &server.Server{}

	server.Start()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
}
