package gigs

import (
	"context"
	"time"

	"github.com/bochkov/m17go/internal/place"
)

type Gig struct {
	Id          int       `json:"id" db:"id"`
	DateTime    time.Time `json:"datetime" db:"datetime"`
	Place       int       `db:"place_id"`
	Description string    `json:"description" db:"description"`
	Url         string    `json:"url" db:"url"`
}

type RsGig struct {
	Id          int         `json:"id"`
	DateTime    time.Time   `json:"datetime"`
	Description string      `json:"description"`
	Url         string      `json:"url"`
	Place       place.Place `json:"place"`
}

type Repository interface {
	findSince(ctx context.Context, since time.Time) ([]Gig, error)
}

type Service interface {
	gigsFrom(ctx context.Context, since time.Time) ([]RsGig, error)
	AllGigs(ctx context.Context) ([]RsGig, error)
	FutureGigs(ctx context.Context) ([]RsGig, error)
}
