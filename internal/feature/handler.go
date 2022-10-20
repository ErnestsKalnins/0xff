package feature

import (
	"encoding/json"
	"github.com/ErnestsKalnins/0xff/pkg/render"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rs/zerolog/hlog"
	"net/http"
)

func NewHandler(service Service) Handler {
	return Handler{svc: service}
}

type Handler struct {
	svc Service
}

func (h Handler) ListFeatures(w http.ResponseWriter, r *http.Request) {
	fs, err := h.svc.store.findAllFeatures(r.Context())
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

	f, err := h.svc.store.findFeature(r.Context(), id)
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

	if err := h.svc.store.saveFeature(r.Context(), feature{
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

	if err := h.svc.store.deleteFeature(r.Context(), id); err != nil {
		hlog.FromRequest(r).
			Error().
			Err(err).
			Msg("delete feature")
		render.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
