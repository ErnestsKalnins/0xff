package sqlite

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ErnestsKalnins/0xff/feature"
)

func TestProjects(t *testing.T) {
	var (
		store = newProjectStore(setup(t))

		id   = uuid.New()
		now  = time.Now().Unix()
		want = feature.Project{
			ID:        id,
			Name:      "Test Project",
			CreatedAt: now,
			UpdatedAt: now,
		}

		ctx = context.Background()
	)

	if err := store.SaveProject(ctx, want); err != nil {
		t.Fatalf("failed to save project: %s", err)
	}

	got, err := store.FindProject(ctx, id)
	assert.Equal(t, want, *got)
	assert.Nil(t, err)

	gotAll, err := store.FindAllProjects(ctx)
	assert.Equal(t, []feature.Project{want}, gotAll)
	assert.Nil(t, err)

	err = store.DeleteProject(ctx, id)
	assert.Nil(t, err)

	got, err = store.FindProject(ctx, id)
	assert.Nil(t, got)
	assert.Equal(t, err, feature.ErrProjectNotFound{ID: id})
}
