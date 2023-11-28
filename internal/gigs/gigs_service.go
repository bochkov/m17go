package gigs

import (
	"context"
	"log"
	"time"

	"github.com/bochkov/m17go/internal/place"
)

type service struct {
	gigsRepo   Repository
	placesRepo place.Repository
	timeout    time.Duration
}

func NewService(gigs Repository, places place.Repository) Service {
	return &service{
		gigs,
		places,
		time.Duration(2) * time.Second,
	}
}

func (s *service) AllGigs(ctx context.Context) ([]RsGig, error) {
	since := time.Date(1982, 10, 16, 18, 0, 0, 0, time.Local)
	return s.gigsFrom(ctx, since)
}

func (s *service) FutureGigs(ctx context.Context) ([]RsGig, error) {
	since := time.Now().Truncate(24 * time.Hour)
	return s.gigsFrom(ctx, since)
}

func (s *service) gigsFrom(c context.Context, since time.Time) ([]RsGig, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	gigs, err := s.gigsRepo.findSince(ctx, since)
	if err != nil {
		return nil, err
	}

	result := make([]RsGig, 0)
	for _, gig := range gigs {
		place, err := s.placesRepo.FindById(ctx, gig.Place)
		if err != nil {
			log.Fatalf("%v", err.Error())
		} else {
			dt := time.Date(
				gig.Date.Year(), gig.Date.Month(), gig.Date.Day(),
				gig.Time.Hour(), gig.Time.Minute(), gig.Time.Second(),
				0, gig.Time.Location())
			rg := &RsGig{
				Id:    gig.Id,
				Date:  dt,
				Desc:  gig.Desc,
				Url:   gig.Url,
				Place: *place,
			}
			result = append(result, *rg)
		}
	}
	return result, nil
}
