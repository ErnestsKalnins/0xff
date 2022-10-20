package feature

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
)

func NewStore(db *sql.DB) Store {
	return Store{db: db}
}

type Store struct {
	db *sql.DB
}

func (s Store) findAllFeatures(ctx context.Context) ([]feature, error) {
	rs, err := s.db.QueryContext(
		ctx,
		`SELECT id, technical_name, display_name, description, enabled, created_at, updated_at FROM features`,
	)
	if err != nil {
		return nil, err
	}

	var fs []feature
	for rs.Next() {
		var f feature
		if err := rs.Scan(
			&f.ID,
			&f.TechnicalName,
			&f.DisplayName,
			&f.Description,
			&f.Enabled,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}

	if err := rs.Err(); err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}
	return fs, nil
}

func (s Store) findFeature(ctx context.Context, id uuid.UUID) (*feature, error) {
	f := feature{ID: id}
	if err := s.db.QueryRowContext(
		ctx,
		`SELECT technical_name, display_name, description, enabled, created_at, updated_at FROM features WHERE id = ?`,
	).Scan(
		&f.TechnicalName,
		&f.DisplayName,
		&f.Description,
		&f.Enabled,
		&f.CreatedAt,
		&f.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &f, nil
}

func (s Store) saveFeature(ctx context.Context, f feature) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO features(technical_name,display_name,description,enabled) VALUES(?,?,?,?)`,
		f.TechnicalName, f.DisplayName, f.Description, f.Enabled,
	)
	return err
}

func (s Store) deleteFeature(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(
		ctx,
		`DELETE FROM features WHERE id = ?`,
		id,
	)
	return err
}
