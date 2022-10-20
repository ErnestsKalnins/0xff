package feature

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func NewStore(db *sql.DB) Store {
	return Store{db: db}
}

type Store struct {
	db *sql.DB
}

func (s Store) findAllProjectFeatures(ctx context.Context, projectID uuid.UUID) ([]feature, error) {
	rs, err := s.db.QueryContext(
		ctx,
		`SELECT id, technical_name, display_name, description, created_at, updated_at FROM features WHERE project_id = ?`,
		projectID,
	)
	if err != nil {
		return nil, err
	}

	var fs []feature
	for rs.Next() {
		f := feature{ProjectID: projectID}
		if err := rs.Scan(
			&f.ID,
			&f.TechnicalName,
			&f.DisplayName,
			&f.Description,
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

type errNotFound struct {
	id uuid.UUID
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("could not find feature by id %s", e.id)
}

func (e errNotFound) Code() int { return http.StatusNotFound }

func (s Store) findFeature(ctx context.Context, id uuid.UUID) (*feature, error) {
	f := feature{ID: id}
	if err := s.db.QueryRowContext(
		ctx,
		`SELECT technical_name, display_name, description, created_at, updated_at FROM features WHERE id = ?`,
	).Scan(
		&f.TechnicalName,
		&f.DisplayName,
		&f.Description,
		&f.CreatedAt,
		&f.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errNotFound{id: id}
		}
		return nil, err
	}

	return &f, nil
}

func (s Store) saveFeature(ctx context.Context, f feature) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO features (id, technical_name,display_name,description) VALUES (?,?,?,?)`,
		f.ID, f.TechnicalName, f.DisplayName, f.Description,
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
