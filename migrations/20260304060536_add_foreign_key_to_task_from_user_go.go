package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddForeignKeyToTaskFromUserGo, downAddForeignKeyToTaskFromUserGo)
}

func upAddForeignKeyToTaskFromUserGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx,
		`
			ALTER TABLE tasks
			ADD COLUMN user_id UUID;

			ALTER TABLE tasks
			ADD CONSTRAINT fk_user_id
			FOREIGN KEY (user_id) REFERENCES users(id);
		`)
	return err
}

func downAddForeignKeyToTaskFromUserGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		ALTER TABLE tasks DROP CONSTRAINT fk_user_id;

		ALTER TABLE tasks DROP COLUMN user_id;
		`)
	return err
}
