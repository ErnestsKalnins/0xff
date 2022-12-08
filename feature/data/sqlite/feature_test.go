package sqlite

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ErnestsKalnins/0xff/feature"
)

func TestFeatures(t *testing.T) {
	var (
		db    = setup(t)
		store = newFeatureStore(db)

		id        = uuid.New()
		projectID = uuid.New()
		now       = time.Now().Unix()
		want      = feature.Feature{
			ID:            id,
			ProjectID:     projectID,
			TechnicalName: "Test Technical Name",
			DisplayName:   addrOf("Test Display Name"),
			Description:   addrOf("Test Description"),
			CreatedAt:     now,
			UpdatedAt:     now,
		}

		ctx = context.Background()
	)

	if err := newProjectStore(db).SaveProject(ctx, feature.Project{ID: projectID}); err != nil {
		t.Fatalf("failed to save project: %s", err)
	}

	if err := store.SaveFeature(ctx, want); err != nil {
		t.Fatalf("failed to save feature: %s", err)
	}

	got, err := store.FindFeature(ctx, id)
	assert.Equal(t, want, *got)
	assert.Nil(t, err)

	gotAll, err := store.FindAllProjectFeatures(ctx, projectID)
	assert.Equal(t, []feature.Feature{want}, gotAll)
	assert.Nil(t, err)

	err = store.DeleteFeature(ctx, id)
	assert.Nil(t, err)

	got, err = store.FindFeature(ctx, id)
	assert.Nil(t, got)
	assert.Equal(t, err, feature.ErrFeatureNotFound{ID: id})
}

func addrOf[T any](t T) *T {
	return &t
}
