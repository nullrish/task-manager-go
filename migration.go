package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/joho/godotenv"

	_ "github.com/nullrish/task-manager-go/migrations"
)

func main() {
	cmd := "up"

	if len(os.Args) >= 2 {
		cmd = strings.ToLower(os.Args[1])
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatal(err)
	}
	switch cmd {
	case "up":
		if err := goose.Up(db, "./migrations"); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := goose.Down(db, "./migrations"); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Unknown command: %s (up or down)", cmd)
	}
}
