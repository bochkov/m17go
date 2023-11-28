package link

import "context"

type Link struct {
	Id       int    `db:"id"`
	Url      string `db:"url"`
	ProvId   int    `db:"provid"`
	Provider string `db:"provider"`
}

type RsLink struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	ProvId   int    `json:"provid"`
	Provider string `json:"provider"`
}

type Repository interface {
	LinksFor(ctx context.Context, id int) ([]Link, error)
}
