package sqlite

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/ErnestsKalnins/0xff/feature"
	"github.com/google/uuid"
)

func TestFeatureStates(t *testing.T) {
	var (
		db    = setup(t)
		store = newFeatureStateStore(db)

		now = time.Now().Unix()

		project = feature.Project{
			ID:        uuid.New(),
			Name:      "Test Project",
			CreatedAt: now,
			UpdatedAt: now,
		}
		environment = feature.Environment{
			ID:        uuid.New(),
			ProjectID: project.ID,
			Name:      "Test Environment",
			CreatedAt: now,
			UpdatedAt: now,
		}
		feat = feature.Feature{
			ID:            uuid.New(),
			ProjectID:     project.ID,
			TechnicalName: "Test Technical Name",
			DisplayName:   addrOf("Test Display Name"),
			Description:   addrOf("Test Description"),
			CreatedAt:     now,
			UpdatedAt:     now,
		}

		want = feature.EnvironmentFeature{
			FeatureID:     feat.ID,
			TechnicalName: feat.TechnicalName,
			DisplayName:   feat.DisplayName,
			Description:   feat.Description,
			State:         feature.StateConstant(false),
			CreatedAt:     now,
			UpdatedAt:     now,
		}

		ctx = context.Background()
	)

	if err := newProjectStore(db).SaveProject(ctx, project); err != nil {
		t.Fatalf("failed to save project: %s", err)
	}

	if err := newEnvironmentStore(db).SaveEnvironment(ctx, environment); err != nil {
		t.Fatalf("failed to save environment: %s", err)
	}

	if err := newFeatureStore(db).SaveFeature(ctx, feat); err != nil {
		t.Fatalf("failed to save feature: %s", err)
	}

	got, err := store.FindAllEnvironmentFeatures(ctx, environment.ID)
	assert.Equal(t, []feature.EnvironmentFeature{want}, got)
	assert.Nil(t, err)
}
