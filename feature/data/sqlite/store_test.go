package sqlite

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/ErnestsKalnins/0xff/pkg/config"
	_ "github.com/ErnestsKalnins/0xff/pkg/config/test"
	"github.com/ErnestsKalnins/0xff/pkg/migrate"
)

func setup(t *testing.T) *sql.DB {
	t.Helper()
	x := config.DSN()
	_ = x
	db, err := sql.Open("sqlite3", config.DSN())
	if err != nil {
		t.Fatalf("failed to open database connection: %s", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("failed to ping database: %s", err)
	}

	if err := migrate.Migrate(db); err != nil {
		t.Fatalf("failed to run migrations: %s", err)
	}

	return db
}
