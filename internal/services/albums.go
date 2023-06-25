package services

import (
	"database/sql"
	"log"
)

type MType int16

const (
	AlbumType  MType = 1
	SingleType MType = 2
)

type Link struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	ProvId   int    `json:"provid"`
	Provider string `json:"provider"`
}

type Album struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	MType MType  `json:"type"`
	Year  int    `json:"year"`
	Slug  string `json:"slug"`
	Links []Link `json:"links"`
}

type Albums struct {
	db *sql.DB
}

func NewAlbums(db *sql.DB) *Albums {
	return &Albums{db: db}
}

func (albums Albums) LinksFor(id int) []Link {
	queryStr := `SELECT ml.id, mp.id, mp.name, ml.url 
			     FROM music_links ml, music_provs mp
			     WHERE ml.provider = mp.id and ml.music = $1
			     ORDER BY mp.id, ml.id`
	links := make([]Link, 0)
	rows, err := albums.db.Query(queryStr, id)
	if err != nil {
		log.Println(err)
		return links
	}
	defer rows.Close()
	for rows.Next() {
		var l Link
		err := rows.Scan(&l.Id, &l.ProvId, &l.Provider, &l.Url)
		if err != nil {
			return nil
		}
		links = append(links, l)
	}
	return links
}

func (albums Albums) Promo() Album {
	queryStr := `SELECT mus.id, mus.name, mus.year, mus.type, mus.slug 
			     FROM music mus 
			     WHERE mus.id = (SELECT max(mu.id) FROM music mu WHERE mu.ignore=false)`
	var a Album
	albums.db.
		QueryRow(queryStr).
		Scan(&a.Id, &a.Name, &a.Year, &a.MType, &a.Slug)
	a.Links = albums.LinksFor(a.Id)
	return a
}

func (albums Albums) Albums() []Album {
	queryStr := `SELECT mus.id, mus.name, mus.year, mus.type, mus.slug
			     FROM music mus
			     WHERE mus.ignore=false
			     ORDER by mus.year desc, mus.id desc`
	allAlbums := make([]Album, 0)
	rows, err := albums.db.Query(queryStr)
	if err != nil {
		log.Println(err)
		return allAlbums
	}
	defer rows.Close()
	for rows.Next() {
		var a Album
		rows.Scan(&a.Id, &a.Name, &a.Year, &a.MType, &a.Slug)
		a.Links = albums.LinksFor(a.Id)
		allAlbums = append(allAlbums, a)
	}
	return allAlbums
}

func (albums Albums) AlbumsOf(mType MType) []Album {
	var result []Album
	for _, album := range albums.Albums() {
		if album.MType == mType {
			result = append(result, album)
		}
	}
	return result
}
