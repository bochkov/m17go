package services

import (
	"database/sql"
	"log"
	"time"
)

type gig struct {
	Id    int       `json:"id"`
	Date  time.Time `json:"date"`
	Desc  string    `json:"desc"`
	Url   string    `json:"url"`
	Place Place     `json:"place"`
}

type Gigs struct {
	db *sql.DB
}

func NewGigs(db *sql.DB) *Gigs {
	return &Gigs{db: db}
}

func (gigs Gigs) Find(since time.Time) []gig {
	result := make([]gig, 0)
	rows, err := gigs.db.
		Query(`SELECT g.* 
			FROM gigs g 
			WHERE g.dt >= $1 
			ORDER BY g.dt desc, g.tm`, since)
	if err != nil {
		log.Println(err)
		return result
	}
	defer rows.Close()
	for rows.Next() {
		var g gig
		var placeId int
		var dt time.Time
		var tm time.Time
		rows.Scan(&g.Id, &dt, &tm, &placeId, &g.Desc, &g.Url)
		g.Date = time.Date(dt.Year(), dt.Month(), dt.Day(), tm.Hour(), tm.Minute(), tm.Second(), 0, tm.Location())
		g.Place = NewPlaces(gigs.db).FindById(placeId)
		result = append(result, g)
	}
	return result
}
