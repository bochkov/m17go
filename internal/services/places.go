package services

import "database/sql"

type Place struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Link         string `json:"link"`
	Slug         string `json:"slug"`
	InvertedLogo bool   `json:"inverted-logo"`
}

type Places struct {
	db *sql.DB
}

func NewPlaces(db *sql.DB) *Places {
	return &Places{db: db}
}

func (places Places) FindById(id int) (p Place) {
	err := places.db.
		QueryRow("SELECT * FROM places WHERE id = $1", id).
		Scan(&p.Id, &p.Name, &p.Address, &p.Link, &p.Slug, &p.InvertedLogo)
	if err != nil {
		return Place{}
	}
	return
}
