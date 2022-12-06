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

func (h Handler) ListEnvironments(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "projectId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	es, err := h.store.FindAllProjectEnvironments(r.Context(), projectID)
	if err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("list environments")
		render.Error(w, err)
		return
	}

	render.JSON(w, es)
}

func (h Handler) GetEnvironment(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "environmentId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	f, err := h.store.FindEnvironment(r.Context(), id)
	if err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("get environment")
		render.Error(w, err)
		return
	}

	render.JSON(w, f)
}

func (h Handler) SaveEnvironment(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "projectId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.saveEnvironment.Handle(r.Context(), feature.Environment{
		ProjectID: projectID,
		Name:      req.Name,
	}); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("save environment")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) DeleteEnvironment(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "environmentId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.store.DeleteEnvironment(r.Context(), id); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("delete environment")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
