package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitializeDatabase(connString string) *sql.DB {
	log.Println("🗄 Opening database")
	conn, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatal("❌ Failed to initialize database:", err.Error())
	}
	log.Println("✅ Successfully opened database!")
	fmt.Println()
	return conn
}
