package db

import (
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	migrate "github.com/rubenv/sql-migrate"

	"github.com/jmoiron/sqlx"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
