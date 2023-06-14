package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bochkov/m17go/internal/services"
)

type Api struct {
	gigs    *services.Gigs
	members *services.Members
	albums  *services.Albums
}

func ConfigureController(db *sql.DB, mux *http.ServeMux) {
	api := &Api{
		gigs:    services.NewGigs(db),
		members: services.NewMembers(db),
		albums:  services.NewAlbums(db),
	}
	mux.HandleFunc("/api/v1/gigs/all", api.allGigs)
	mux.HandleFunc("/api/v1/gigs", api.futureGigs)
	mux.HandleFunc("/api/v1/members/all", api.allMembers)
	mux.HandleFunc("/api/v1/members", api.actualMembers)
	mux.HandleFunc("/api/v1/albums/all", api.allAlbums)
	mux.HandleFunc("/api/v1/albums", api.onlyAlbums)
	mux.HandleFunc("/api/v1/albums/singles", api.onlySingles)
	mux.HandleFunc("/api/v1/promo", api.promo)
}

func renderJson(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (api *Api) gigsFrom(since time.Time, w http.ResponseWriter) {
	gigs := api.gigs.Find(since)
	renderJson(w, gigs)
}

func (api *Api) allGigs(w http.ResponseWriter, req *http.Request) {
	since := time.Date(1982, 10, 16, 18, 0, 0, 0, time.Local)
	api.gigsFrom(since, w)
}

func (api *Api) futureGigs(w http.ResponseWriter, req *http.Request) {
	since := time.Now().Truncate(24 * time.Hour)
	api.gigsFrom(since, w)
}

func (api *Api) allMembers(w http.ResponseWriter, req *http.Request) {
	members := api.members.FindAll()
	renderJson(w, members)
}

func (api *Api) actualMembers(w http.ResponseWriter, req *http.Request) {
	members := api.members.FindActual()
	renderJson(w, members)
}

func (api *Api) allAlbums(w http.ResponseWriter, req *http.Request) {
	albums := api.albums.Albums()
	renderJson(w, albums)
}

func (api *Api) onlyAlbums(w http.ResponseWriter, req *http.Request) {
	result := api.albums.AlbumsOf(services.ALBUM_TYPE)
	renderJson(w, result)
}

func (api *Api) onlySingles(w http.ResponseWriter, req *http.Request) {
	result := api.albums.AlbumsOf(services.SINGLE_TYPE)
	renderJson(w, result)
}

func (api *Api) promo(w http.ResponseWriter, req *http.Request) {
	album := api.albums.Promo()
	renderJson(w, album)
}
