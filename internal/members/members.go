package members

import (
	"context"
)

type Member struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Actual     bool   `db:"actual"`
	Instrument string `db:"instrument"`
}

type RsMember struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Instrument string `json:"instrument"`
	Actual     bool   `json:"actual"`
}

type Repository interface {
	findAll(ctx context.Context) ([]Member, error)
	findActual(ctx context.Context) ([]Member, error)
}

type Service interface {
	FindAll(ctx context.Context) ([]RsMember, error)
	FindActual(ctx context.Context) ([]RsMember, error)
}
