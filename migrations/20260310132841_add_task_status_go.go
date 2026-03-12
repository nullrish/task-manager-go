package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddTaskStatusGo, downAddTaskStatusGo)
}

func upAddTaskStatusGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downAddTaskStatusGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
