package songs

import "context"

type Song struct {
	Id       int    `json:"id"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	Lyrics   string `json:"lyrics"`
}

type Repository interface {
	FindAllForAlbum(ctx context.Context, albumId int) ([]Song, error)
}
