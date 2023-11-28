package members

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) FindAll(c context.Context) ([]RsMember, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	members, err := s.Repository.findAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]RsMember, 0)
	for _, mem := range members {
		m := &RsMember{
			Id:         mem.Id,
			Name:       mem.Name,
			Instrument: mem.Instrument,
			Actual:     mem.Actual,
		}
		result = append(result, *m)
	}
	return result, nil
}

func (s *service) FindActual(c context.Context) ([]RsMember, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	members, err := s.Repository.findActual(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]RsMember, 0)
	for _, mem := range members {
		m := &RsMember{
			Id:         mem.Id,
			Name:       mem.Name,
			Instrument: mem.Instrument,
			Actual:     mem.Actual,
		}
		result = append(result, *m)
	}
	return result, nil
}
