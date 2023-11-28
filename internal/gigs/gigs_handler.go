package gigs

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

// AllGigs handles "/api/v1/gigs/all"
func (h *Handler) AllGigs(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.AllGigs(r.Context())
	util.DefaultHandle(w, r, res, err)
}

// FutureGigs handles "/api/v1/gigs"
func (h *Handler) FutureGigs(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.FutureGigs(r.Context())
	util.DefaultHandle(w, r, res, err)
}
