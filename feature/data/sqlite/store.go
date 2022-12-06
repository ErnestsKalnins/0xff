package sqlite

import "database/sql"

func New(db *sql.DB) Store {
	return Store{
		projectStore:      newProjectStore(db),
		environmentStore:  newEnvironmentStore(db),
		featureStore:      newFeatureStore(db),
		featureStateStore: newFeatureStateStore(db),
	}
}

type Store struct {
	projectStore
	environmentStore
	featureStore
	featureStateStore
}
