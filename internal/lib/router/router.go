package router

import (
	"context"
	"net/http"
	"time"

	"github.com/bochkov/m17go/internal/albums"
	"github.com/bochkov/m17go/internal/gigs"
	"github.com/bochkov/m17go/internal/members"
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

	r.Route("/api/v1/lyric", func(r chi.Router) {
		r.Route("/{albumSlug:[a-z-]+}", func(r chi.Router) {
			r.Use(AlbumCtx)
			r.Route("/{songSlug:[a-z-]+}", func(r chi.Router) {
				r.Use(SongCtx)
				r.Get("/", albums.OneSongInAlbum)
			})
			r.Get("/", albums.AllSongsInAlbum)
		})
		r.Get("/", albums.AllSongs)
	})

	r.Route("/api/v1/albums", func(r chi.Router) {
		r.Get("/all", albums.AllAlbums)
		r.Get("/singles", albums.OnlySingles)
		r.Get("/", albums.OnlyAlbums)
	})

	r.Route("/api/v1/gigs", func(r chi.Router) {
		r.Get("/all", gigs.AllGigs)
		r.Get("/", gigs.FutureGigs)
	})

	r.Route("/api/v1/members", func(r chi.Router) {
		r.Get("/all", members.AllMembers)
		r.Get("/", members.ActualMembers)
	})

	r.Get("/api/v1/promo", albums.Promo)

	return r
}

func AlbumCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			albumSlug := chi.URLParam(r, "albumSlug")
			ctx := context.WithValue(r.Context(), "albumSlug", albumSlug)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func SongCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			songSlug := chi.URLParam(r, "songSlug")
			ctx := context.WithValue(r.Context(), "songSlug", songSlug)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
