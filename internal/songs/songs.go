package songs

import "context"

type Song struct {
	Id       int    `json:"id"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Lyrics   string `json:"lyrics"`
}

type SongWithAlbum struct {
	Song
	AlbumName string `json:"album_name"`
	AlbumYear int    `json:"album_year"`
	AlbumSlug string `json:"album_slug"`
}

type Repository interface {
	AllSongs(ctx context.Context) ([]SongWithAlbum, error)
	SongsForAlbum(ctx context.Context, albumSlug string) ([]Song, error)
}
