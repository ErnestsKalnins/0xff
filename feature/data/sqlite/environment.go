package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func newEnvironmentStore(db *sql.DB) environmentStore {
	return environmentStore{db: db}
}

type environmentStore struct {
	db *sql.DB
}

func (s environmentStore) FindAllProjectEnvironments(ctx context.Context, projectID uuid.UUID) ([]feature.Environment, error) {
	rs, err := s.db.QueryContext(
		ctx,
		`SELECT id, name, created_at, updated_at FROM environments WHERE project_id = ?`,
		projectID,
	)
	if err != nil {
		return nil, err
	}

	var es []feature.Environment
	for rs.Next() {
		e := feature.Environment{ProjectID: projectID}
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

func (s environmentStore) FindEnvironment(ctx context.Context, id uuid.UUID) (*feature.Environment, error) {
	e := feature.Environment{ID: id}
	if err := s.db.QueryRowContext(
		ctx,
		`SELECT project_id, name, created_at, updated_at FROM environments WHERE id = ?`,
		id,
	).Scan(
		&e.ProjectID,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, feature.ErrEnvironmentNotFound{ID: id}
		}
		return nil, err
	}

	return &e, nil
}

func (s environmentStore) SaveEnvironment(ctx context.Context, e feature.Environment) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO environments (id, project_id, name, created_at, updated_at) VALUES (?,?,?,?,?)`,
		e.ID, e.ProjectID, e.Name, e.CreatedAt, e.UpdatedAt,
	)
	return err
}

func (s environmentStore) DeleteEnvironment(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(
		ctx,
		`DELETE FROM environments WHERE id = ?`,
		id,
	)
	return err
}
