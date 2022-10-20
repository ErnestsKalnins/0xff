package project

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

func (s Store) findAllProjects(ctx context.Context) ([]project, error) {
	rs, err := s.db.QueryContext(
		ctx,
		`SELECT id, name FROM projects`,
	)
	if err != nil {
		return nil, err
	}

	var ps []project
	for rs.Next() {
		var p project
		if err := rs.Scan(
			&p.ID,
			&p.Name,
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

type errNotFound struct {
	id uuid.UUID
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("could not find project by id %s", e.id)
}

func (e errNotFound) Code() int {
	return http.StatusNotFound
}

func (s Store) findProject(ctx context.Context, id uuid.UUID) (*project, error) {
	p := project{ID: id}
	if err := s.db.QueryRowContext(
		ctx,
		`SELECT name FROM projects WHERE id = ?`,
	).Scan(&p.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errNotFound{id: id}
		}
		return nil, err
	}

	return &p, nil
}

func (s Store) saveProject(ctx context.Context, p project) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO projects (id, name) VALUES(?,?)`,
		p.ID, p.Name,
	)
	return err
}

func (s Store) deleteProject(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(
		ctx,
		`DELETE FROM projects WHERE id = ?`,
		id,
	)
	return err
}
