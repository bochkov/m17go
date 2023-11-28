package place

import "context"

type Place struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Address      string `json:"address" db:"address"`
	Link         string `json:"link" db:"link"`
	Slug         string `json:"slug" db:"slug"`
	InvertedLogo bool   `json:"inverted-logo" db:"inverted_logo"`
}

type Repository interface {
	FindById(ctx context.Context, id int) (*Place, error)
}

type Service interface {
}
