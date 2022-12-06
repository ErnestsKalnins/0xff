package api

import (
	"github.com/ErnestsKalnins/0xff/feature/usecase"
)

func NewHandler(store usecase.Store) Handler {
	return Handler{
		store: store,

		saveProject:      usecase.NewSaveProjectHandler(store),
		saveEnvironment:  usecase.NewSaveEnvironmentHandler(store),
		saveFeature:      usecase.NewSaveFeatureHandler(store),
		saveFeatureState: usecase.NewSaveFeatureStateHandler(store),
	}
}

type Handler struct {
	store usecase.Store

	saveProject      usecase.SaveProjectHandler
	saveEnvironment  usecase.SaveEnvironmentHandler
	saveFeature      usecase.SaveFeatureHandler
	saveFeatureState usecase.SaveFeatureStateHandler
}
