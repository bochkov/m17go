package albums

import (
	"net/http"

	"github.com/bochkov/m17go/internal/util"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{Service: s}
}

// AllAlbums handles "/api/v1/albums/all"
func (h *Handler) AllAlbums(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.AllAlbums(r.Context())
	util.DefaultHandle(w, r, res, err)
}

// OnlyAlbums handles "/api/v1/albums"
func (h *Handler) OnlyAlbums(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.OnlyAlbums(r.Context())
	util.DefaultHandle(w, r, res, err)
}

// OnlySingles handles "/api/v1/albums/singles"
func (h *Handler) OnlySingles(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.OnlySingles(r.Context())
	util.DefaultHandle(w, r, res, err)
}

// Promo handles "/api/v1/promo"
func (h *Handler) Promo(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.Promo(r.Context())
	util.DefaultHandle(w, r, res, err)
}
