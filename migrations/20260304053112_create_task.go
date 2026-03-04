package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTask, downCreateTask)
}

func upCreateTask(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx,
		`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

			CREATE TABLE tasks (
				id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
				task_title TEXT NOT NULL,
				task_description TEXT,
				created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
			);
		`)
	return err
}

func downCreateTask(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, "DROP TABLE IF NOT EXISTS tasks;")
	return err
}
