package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func newFeatureStore(db *sql.DB) featureStore {
	return featureStore{db: db}
}

type featureStore struct {
	db *sql.DB
}

func (s featureStore) FindAllProjectFeatures(ctx context.Context, projectID uuid.UUID) ([]feature.Feature, error) {
	rs, err := s.db.QueryContext(
		ctx,
		`SELECT id, technical_name, display_name, description, created_at, updated_at FROM features WHERE project_id = ?`,
		projectID,
	)
	if err != nil {
		return nil, err
	}

	var fs []feature.Feature
	for rs.Next() {
		f := feature.Feature{ProjectID: projectID}
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

func (s featureStore) FindFeature(ctx context.Context, id uuid.UUID) (*feature.Feature, error) {
	f := feature.Feature{ID: id}
	if err := s.db.QueryRowContext(
		ctx,
		`SELECT project_id, technical_name, display_name, description, created_at, updated_at FROM features WHERE id = ?`,
		id,
	).Scan(
		&f.ProjectID,
		&f.TechnicalName,
		&f.DisplayName,
		&f.Description,
		&f.CreatedAt,
		&f.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, feature.ErrFeatureNotFound{ID: id}
		}
		return nil, err
	}

	return &f, nil
}

func (s featureStore) SaveFeature(ctx context.Context, f feature.Feature) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO features (id, project_id, technical_name, display_name, description, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`,
		f.ID, f.ProjectID, f.TechnicalName, f.DisplayName, f.Description, f.CreatedAt, f.UpdatedAt,
	)
	return err
}

func (s featureStore) DeleteFeature(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(
		ctx,
		`DELETE FROM features WHERE id = ?`,
		id,
	)
	return err
}
