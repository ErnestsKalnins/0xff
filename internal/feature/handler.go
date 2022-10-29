package feature

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rs/zerolog/hlog"

	"github.com/ErnestsKalnins/0xff/pkg/render"
)

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

type Handler struct {
	service Service
}

func (h Handler) ListFeatures(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "projectId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	fs, err := h.service.store.findAllProjectFeatures(r.Context(), projectID)
	if err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("list features")
		render.Error(w, err)
		return
	}

	render.JSON(w, fs)
}

func (h Handler) GetFeature(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "featureId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	f, err := h.service.store.findFeature(r.Context(), id)
	if err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("get feature")
		render.Error(w, err)
		return
	}

	render.JSON(w, f)
}

func (h Handler) SaveFeature(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "projectId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	var req struct {
		TechnicalName string  `json:"technicalName"`
		DisplayName   *string `json:"displayName"`
		Description   *string `json:"description"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.service.saveFeature(r.Context(), feature{
		ProjectID:     projectID,
		TechnicalName: req.TechnicalName,
		DisplayName:   req.DisplayName,
		Description:   req.Description,
	}); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("save feature")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) DeleteFeature(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "featureId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.service.store.deleteFeature(r.Context(), id); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("delete feature")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) ListFeatureStates(w http.ResponseWriter, r *http.Request) {
	environmentID, err := uuid.Parse(chi.URLParam(r, "environmentId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	ss, err := h.service.store.findAllEnvironmentFeatures(r.Context(), environmentID)
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
		State featureStateTransport `json:"state"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.service.setFeatureState(
		r.Context(),
		environmentID,
		featureID,
		req.State.value,
	); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("set feature state")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
