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

func (h Handler) ListProjects(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.FindAllProjects(r.Context())
	if err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("list projects")
		render.Error(w, err)
		return
	}

	render.JSON(w, ps)
}

func (h Handler) GetProject(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "projectId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	p, err := h.store.FindProject(r.Context(), id)
	if err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("get project")
		render.Error(w, err)
		return
	}

	render.JSON(w, p)
}

func (h Handler) SaveProject(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.saveProject.Handle(r.Context(), feature.Project{Name: req.Name}); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("save project")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "projectId"))
	if err != nil {
		render.Error(w, render.TagBadRequest(err))
		return
	}

	if err := h.store.DeleteProject(r.Context(), id); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("delete project")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
