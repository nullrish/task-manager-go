package main

import (
	"github.com/nullrish/task-manager-go/internal/app"
	"github.com/nullrish/task-manager-go/internal/db"
)

// Entry point to the server
func main() {
	db := db.InitializeDatabase()
	s := app.InitializeServer("127.0.0.1", "8080", db)
	s.StartServer()
}
