package sqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func newFeatureStateStore(db *sql.DB) featureStateStore {
	return featureStateStore{db: db}
}

type featureStateStore struct {
	db *sql.DB
}

func (s featureStateStore) FindAllEnvironmentFeatures(ctx context.Context, environmentID uuid.UUID) ([]feature.EnvironmentFeature, error) {
	//language=sqlite
	const findAllEnvironmentFeaturesQuery = `
	WITH environment_feature_states AS
	    (SELECT * FROM feature_states WHERE environment_id = ?)
	SELECT
	    f.id,
		f.technical_name,
		f.display_name,
		f.description,
		efs.state,
		f.created_at,
		COALESCE(efs.updated_at, f.created_at) as updated_at
	FROM features f
		LEFT JOIN environment_feature_states efs on f.id = efs.feature_id
`

	rs, err := s.db.QueryContext(
		ctx,
		findAllEnvironmentFeaturesQuery,
		environmentID,
	)
	if err != nil {
		return nil, err
	}

	var efs []feature.EnvironmentFeature
	for rs.Next() {
		var (
			ef feature.EnvironmentFeature
			sc stateColumn
		)
		if err := rs.Scan(
			&ef.FeatureID,
			&ef.TechnicalName,
			&ef.DisplayName,
			&ef.Description,
			&sc,
			&ef.CreatedAt,
			&ef.UpdatedAt,
		); err != nil {
			return nil, err
		}
		ef.State = sc.value
		efs = append(efs, ef)
	}

	if err := rs.Err(); err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return efs, nil
}

func (s featureStateStore) SaveFeatureState(ctx context.Context, ef feature.EnvironmentFeature) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO feature_states (id, feature_id, environment_id, state, updated_at) VALUES (?,?,?,?,?,?)`,
		ef.ID, ef.FeatureID, ef.EnvironmentID, stateColumn{value: ef.State}, ef.UpdatedAt,
	)
	return err
}

type stateColumn struct {
	value feature.State
}

// Value implements driver.Valuer.
func (c stateColumn) Value() (driver.Value, error) {
	return feature.StateMarshaler{Value: c.value}.MarshalJSON()
}

// Scan implements sql.Scanner.
func (c *stateColumn) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		c.value = feature.StateConstant(false)
		return nil
	case []byte:
		var sm feature.StateMarshaler
		if err := sm.UnmarshalJSON(v); err != nil {
			return err
		}
		c.value = sm.Value
		return nil
	default:
		return fmt.Errorf("bad featureState type: %T", v)
	}
}
