package services

import (
	"database/sql"
	"log"
)

type MType int16

const (
	ALBUM_TYPE  MType = 1
	SINGLE_TYPE MType = 2
)

type link struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	ProvId   int    `json:"provid"`
	Provider string `json:"provider"`
}

type album struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	MType MType  `json:"type"`
	Year  int    `json:"year"`
	Slug  string `json:"slug"`
	Links []link `json:"links"`
}

type Albums struct {
	db *sql.DB
}

func NewAlbums(db *sql.DB) *Albums {
	return &Albums{db: db}
}

func (albums Albums) LinksFor(id int) []link {
	links := make([]link, 0)
	rows, err := albums.db.
		Query(`SELECT ml.id, mp.id, mp.name, ml.url 
			FROM music_links ml, music_provs mp
			WHERE ml.provider = mp.id and ml.music = $1
			ORDER BY mp.id, ml.id`, id)
	if err != nil {
		log.Println(err)
		return links
	}
	defer rows.Close()

	for rows.Next() {
		var l link
		err := rows.Scan(&l.Id, &l.ProvId, &l.Provider, &l.Url)
		if err != nil {
			return nil
		}
		links = append(links, l)
	}
	return links
}

func (albums Albums) Promo() album {
	var a album
	albums.db.
		QueryRow(`SELECT m.id, m.name, m.year, m.type, m.slug 
			FROM music m WHERE m.id = (SELECT max(id) FROM music WHERE ignore=false)`).
		Scan(&a.Id, &a.Name, &a.Year, &a.MType, &a.Slug)
	a.Links = albums.LinksFor(a.Id)
	return a
}

func (albums Albums) Albums() []album {
	allAlbums := make([]album, 0)
	rows, err := albums.db.
		Query(`SELECT m.id, m.name, m.year, m.type, m.slug
			FROM music m
			WHERE ignore=false
			ORDER by year desc, id desc`)
	if err != nil {
		log.Println(err)
		return allAlbums
	}

	defer rows.Close()
	for rows.Next() {
		var a album
		rows.Scan(&a.Id, &a.Name, &a.Year, &a.MType, &a.Slug)
		a.Links = albums.LinksFor(a.Id)
		allAlbums = append(allAlbums, a)
	}
	return allAlbums
}

func (albums Albums) AlbumsOf(mType MType) []album {
	var result []album
	for _, album := range albums.Albums() {
		if album.MType == mType {
			result = append(result, album)
		}
	}
	return result
}
