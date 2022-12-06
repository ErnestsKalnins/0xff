package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func newProjectStore(db *sql.DB) projectStore {
	return projectStore{db: db}
}

type projectStore struct {
	db *sql.DB
}

func (s projectStore) FindAllProjects(ctx context.Context) ([]feature.Project, error) {
	rs, err := s.db.QueryContext(
		ctx,
		`SELECT id, name, created_at, updated_at FROM projects`,
	)
	if err != nil {
		return nil, err
	}

	var ps []feature.Project
	for rs.Next() {
		var p feature.Project
		if err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}

	if err := rs.Err(); err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}
	return ps, nil
}

func (s projectStore) FindProject(ctx context.Context, id uuid.UUID) (*feature.Project, error) {
	p := feature.Project{ID: id}
	if err := s.db.QueryRowContext(
		ctx,
		`SELECT name, created_at, updated_at FROM projects WHERE id = ?`,
		id,
	).Scan(
		&p.Name,
		&p.CreatedAt,
		&p.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, feature.ErrProjectNotFound{ID: id}
		}
		return nil, err
	}

	return &p, nil
}

func (s projectStore) SaveProject(ctx context.Context, p feature.Project) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO projects (id, name, created_at, updated_at) VALUES(?,?,?,?)`,
		p.ID, p.Name, p.CreatedAt, p.UpdatedAt,
	)
	return err
}

func (s projectStore) DeleteProject(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(
		ctx,
		`DELETE FROM projects WHERE id = ?`,
		id,
	)
	return err
}
