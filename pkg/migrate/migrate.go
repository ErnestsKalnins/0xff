package migrate

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
)

//go:embed migrations
var migrations embed.FS

// Migrate runs all the migrations on the given *sql.DB.
func Migrate(db *sql.DB) error {
	return fs.WalkDir(
		migrations,
		"migrations",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			b, err := migrations.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read file: %w", err)
			}

			if _, err := db.Exec(string(b)); err != nil {
				return fmt.Errorf("execute %q migration: %w", path, err)
			}
			return nil
		},
	)
}
