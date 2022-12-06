package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

type Store interface {
	FindAllProjects(context.Context) ([]feature.Project, error)
	FindProject(context.Context, uuid.UUID) (*feature.Project, error)
	SaveProject(context.Context, feature.Project) error
	DeleteProject(context.Context, uuid.UUID) error

	FindAllProjectEnvironments(context.Context, uuid.UUID) ([]feature.Environment, error)
	FindEnvironment(context.Context, uuid.UUID) (*feature.Environment, error)
	SaveEnvironment(context.Context, feature.Environment) error
	DeleteEnvironment(context.Context, uuid.UUID) error

	FindAllProjectFeatures(context.Context, uuid.UUID) ([]feature.Feature, error)
	FindFeature(context.Context, uuid.UUID) (*feature.Feature, error)
	SaveFeature(context.Context, feature.Feature) error
	DeleteFeature(context.Context, uuid.UUID) error

	FindAllEnvironmentFeatures(context.Context, uuid.UUID) ([]feature.EnvironmentFeature, error)
	SaveFeatureState(context.Context, feature.EnvironmentFeature) error
}
