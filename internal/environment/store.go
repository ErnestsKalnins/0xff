package environment

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

func (s Store) findAllProjectEnvironments(ctx context.Context, projectID uuid.UUID) ([]environment, error) {
	rs, err := s.db.QueryContext(
		ctx,
		`SELECT id, name, created_at, updated_at FROM environments WHERE project_id = ?`,
		projectID,
	)
	if err != nil {
		return nil, err
	}

	var es []environment
	for rs.Next() {
		e := environment{ProjectID: projectID}
		if err := rs.Scan(
			&e.ID,
			&e.Name,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			return nil, err
		}
		es = append(es, e)
	}

	if err := rs.Err(); err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}
	return es, nil
}

type errNotFound struct {
	id uuid.UUID
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("could not find environment by id %s", e.id)
}

func (e errNotFound) Code() int { return http.StatusNotFound }

func (s Store) findEnvironment(ctx context.Context, id uuid.UUID) (*environment, error) {
	e := environment{ID: id}
	if err := s.db.QueryRowContext(
		ctx,
		`SELECT project_id, name, created_at, updated_at FROM features WHERE id = ?`,
	).Scan(
		&e.ProjectID,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errNotFound{id: id}
		}
		return nil, err
	}

	return &e, nil
}

func (s Store) saveEnvironment(ctx context.Context, f environment) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO environments (id, project_id, name) VALUES (?,?,?)`,
		f.ID, f.ProjectID, f.Name,
	)
	return err
}

func (s Store) deleteEnvironment(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(
		ctx,
		`DELETE FROM environments WHERE id = ?`,
		id,
	)
	return err
}
