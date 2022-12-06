package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rs/zerolog/hlog"

	"github.com/ErnestsKalnins/0xff/feature"
	"github.com/ErnestsKalnins/0xff/pkg/render"
)

func (h Handler) ListFeatures(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "projectId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	fs, err := h.store.FindAllProjectFeatures(r.Context(), projectID)
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

	f, err := h.store.FindFeature(r.Context(), id)
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

	if err := h.saveFeature.Handle(r.Context(), feature.Feature{
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

	if err := h.store.DeleteFeature(r.Context(), id); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("delete feature")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
