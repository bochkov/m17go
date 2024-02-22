package router

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bochkov/m17go/internal/albums"
	"github.com/bochkov/m17go/internal/gigs"
	"github.com/bochkov/m17go/internal/members"
	"github.com/bochkov/m17go/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var r *chi.Mux

func InitRouter(albums *albums.Handler, gigs *gigs.Handler, members *members.Handler) *chi.Mux {
	r = chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/api/v1/promo", albums.Promo)
	r.Get("/api/v1/albums/all", albums.AllAlbums)
	r.Get("/api/v1/albums", albums.OnlyAlbums)
	r.Get("/api/v1/albums/singles", albums.OnlySingles)
	r.Route("/api/v1/albums/{albumId:\\d+}/songs", func(r chi.Router) {
		r.Use(AlbumCtx)
		r.Get("/", albums.Songs)
	})

	r.Get("/api/v1/gigs", gigs.FutureGigs)
	r.Get("/api/v1/gigs/all", gigs.AllGigs)

	r.Get("/api/v1/members", members.ActualMembers)
	r.Get("/api/v1/members/all", members.AllMembers)

	return r
}

func AlbumCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "albumId")
			albumId, err := strconv.Atoi(id)
			if err != nil {
				util.DefaultHandle(w, r, nil, err)
			}
			ctx := context.WithValue(r.Context(), "albumId", albumId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
