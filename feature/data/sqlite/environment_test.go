package sqlite

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ErnestsKalnins/0xff/feature"
)

func TestEnvironments(t *testing.T) {
	var (
		db    = setup(t)
		store = newEnvironmentStore(db)

		projectID = uuid.New()
		id        = uuid.New()
		now       = time.Now().Unix()
		want      = feature.Environment{
			ID:        id,
			ProjectID: projectID,
			Name:      "Test Environment",
			CreatedAt: now,
			UpdatedAt: now,
		}

		ctx = context.Background()
	)

	if err := newProjectStore(db).SaveProject(ctx, feature.Project{ID: projectID}); err != nil {
		t.Fatalf("failed to save project: %s", err)
	}

	if err := store.SaveEnvironment(ctx, want); err != nil {
		t.Fatalf("failed to save environment: %s", err)
	}

	got, err := store.FindEnvironment(ctx, id)
	assert.Equal(t, want, *got)
	assert.Nil(t, err)

	gotAll, err := store.FindAllProjectEnvironments(ctx, projectID)
	assert.Equal(t, []feature.Environment{want}, gotAll)
	assert.Nil(t, err)

	err = store.DeleteEnvironment(ctx, id)
	assert.Nil(t, err)

	got, err = store.FindEnvironment(ctx, id)
	assert.Nil(t, got)
	assert.Equal(t, err, feature.ErrEnvironmentNotFound{ID: id})
}
