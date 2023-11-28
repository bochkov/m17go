package members

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

// AllGigs handles "/api/v1/members/all"
func (h *Handler) AllMembers(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.FindAll(r.Context())
	util.DefaultHandle(w, r, res, err)
}

// FutureGigs handles "/api/v1/members"
func (h *Handler) ActualMembers(w http.ResponseWriter, r *http.Request) {
	res, err := h.Service.FindActual(r.Context())
	util.DefaultHandle(w, r, res, err)
}
