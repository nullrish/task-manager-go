package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nullrish/task-manager-go/internal/app"
)

// Entry point to the server
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	s := app.InitializeServer("127.0.0.1", "8080")
	s.StartServer()
}
