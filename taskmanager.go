package main

import (
	"log"

	"github.com/nullrish/task-manager-go/cmd"
)

func main() {
	db := cmd.InitializeDatabase()
	s := cmd.InitializeServer("127.0.0.1", "8080", db)
	log.Fatal(s.StartServer())
}
