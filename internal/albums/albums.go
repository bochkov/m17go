package albums

import (
	"context"

	"github.com/bochkov/m17go/internal/link"
	"github.com/bochkov/m17go/internal/songs"
)

type Album struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Year  int    `db:"year"`
	MType string `db:"type"`
	Slug  string `db:"slug"`
}

type RsAlbum struct {
	Id    int           `json:"id"`
	Name  string        `json:"name"`
	MType string        `json:"type"`
	Year  int           `json:"year"`
	Slug  string        `json:"slug"`
	Links []link.RsLink `json:"links"`
}

type Repository interface {
	FindLatest(ctx context.Context) (*Album, error)
	FindAll(ctx context.Context) ([]Album, error)
}

type Service interface {
	convertLinks(ctx context.Context, albumId int) ([]link.RsLink, error)
	albumsOf(ctx context.Context, mType string) ([]RsAlbum, error)

	AllAlbums(ctx context.Context) ([]RsAlbum, error)
	OnlyAlbums(ctx context.Context) ([]RsAlbum, error)
	OnlySingles(ctx context.Context) ([]RsAlbum, error)
	Promo(ctx context.Context) (*RsAlbum, error)

	AllSongs(ctx context.Context) ([]songs.SongWithAlbum, error)
	SongsInAlbum(ctx context.Context, albumSlug string) ([]songs.Song, error)
	SongInAlbum(ctx context.Context, albumSlug string, songSlug string) (*songs.Song, error)
}
