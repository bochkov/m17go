package gigs

import (
	"context"
	"time"

	"github.com/bochkov/m17go/internal/place"
)

type Gig struct {
	Id    int       `json:"id" db:"id"`
	Date  time.Time `json:"date" db:"dt"`
	Time  time.Time `json:"time" db:"tm"`
	Place int       `db:"place"`
	Desc  string    `json:"desc" db:"desc"`
	Url   string    `json:"url" db:"url"`
}

type RsGig struct {
	Id    int         `json:"id"`
	Date  time.Time   `json:"date"`
	Desc  string      `json:"desc"`
	Url   string      `json:"url"`
	Place place.Place `json:"place"`
}

type Repository interface {
	findSince(ctx context.Context, since time.Time) ([]Gig, error)
}

type Service interface {
	gigsFrom(ctx context.Context, since time.Time) ([]RsGig, error)
	AllGigs(ctx context.Context) ([]RsGig, error)
	FutureGigs(ctx context.Context) ([]RsGig, error)
}
