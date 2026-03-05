package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitializeDatabase(connString string) *sql.DB {
	log.Println("🗄 Opening database")
	db, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatal("❌ Failed to initialize database:", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			err = ctx.Err()
		}
		log.Fatal("❌ Failed to open database: ", err.Error())
	}
	log.Println("✅ Successfully opened database!")
	fmt.Println()
	return db
}
