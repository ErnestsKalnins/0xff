package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rs/zerolog/hlog"

	"github.com/ErnestsKalnins/0xff/feature"
	"github.com/ErnestsKalnins/0xff/feature/usecase"
	"github.com/ErnestsKalnins/0xff/pkg/render"
)

func (h Handler) ListFeatureStates(w http.ResponseWriter, r *http.Request) {
	environmentID, err := uuid.Parse(chi.URLParam(r, "environmentId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	ss, err := h.store.FindAllEnvironmentFeatures(r.Context(), environmentID)
	if err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("list feature states")
		return
	}

	render.JSON(w, ss)
}

func (h Handler) SetFeatureState(w http.ResponseWriter, r *http.Request) {
	environmentID, err := uuid.Parse(chi.URLParam(r, "environmentId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	featureID, err := uuid.Parse(chi.URLParam(r, "featureID"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	var req struct {
		State feature.StateMarshaler `json:"state"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.saveFeatureState.Handle(r.Context(), usecase.SaveFeatureStateRequest{
		EnvironmentID: environmentID,
		FeatureID:     featureID,
		State:         req.State.Value,
	}); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("set feature state")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
