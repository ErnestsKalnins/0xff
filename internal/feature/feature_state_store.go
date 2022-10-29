package feature

import (
	"context"
	"github.com/google/uuid"
)

func (s Store) findAllEnvironmentFeatures(ctx context.Context, environmentID uuid.UUID) ([]readState, error) {
	//language=sqlite
	const findAllEnvironmentFeaturesQuery = `
	WITH environment_feature_states AS
	    (SELECT * FROM feature_states WHERE environment_id = ?)
	SELECT
	    f.id,
		f.technical_name,
		f.display_name,
		f.description,
		eft.state,
		f.created_at,
		COALESCE(eft.updated_at, f.created_at) as updated_at
	FROM features f
		LEFT JOIN environment_feature_states eft on f.id = eft.feature_id
`

	rs, err := s.db.QueryContext(
		ctx,
		findAllEnvironmentFeaturesQuery,
		environmentID,
	)
	if err != nil {
		return nil, err
	}

	var rss []readState
	for rs.Next() {
		var rst readState
		if err := rs.Scan(
			&rst.ID,
			&rst.TechnicalName,
			&rst.DisplayName,
			&rst.Description,
			&rst.State,
			&rst.CreatedAt,
			&rst.UpdatedAt,
		); err != nil {
			return nil, err
		}
		rss = append(rss, rst)
	}

	if err := rs.Err(); err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return rss, nil
}

func (s Store) saveFeatureState(ctx context.Context, ws writeState) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO feature_states (id, feature_id, environment_id, state, updated_at) VALUES (?,?,?,?,?,?)`,
		ws.ID, ws.FeatureID, ws.EnvironmentID, featureStateTransport{value: ws.State}, ws.UpdatedAt,
	)
	return err
}
