package albums

import (
	"context"
	"log"
	"time"

	"github.com/bochkov/m17go/internal/link"
	"github.com/bochkov/m17go/internal/songs"
)

type service struct {
	albums  Repository
	links   link.Repository
	songs   songs.Repository
	timeout time.Duration
}

func NewService(albums Repository, links link.Repository, songs songs.Repository) Service {
	return &service{
		albums,
		links,
		songs,
		time.Duration(2) * time.Second,
	}
}

func (s *service) Promo(c context.Context) (*RsAlbum, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	album, err := s.albums.FindLatest(ctx)
	if err != nil {
		return nil, err
	}

	rlinks, err := s.convertLinks(ctx, album.Id)
	if err != nil {
		return nil, err
	}

	rs := &RsAlbum{
		Id:    album.Id,
		Name:  album.Name,
		MType: album.MType,
		Year:  album.Year,
		Slug:  album.Slug,
		Links: rlinks,
	}
	return rs, nil
}

func (s *service) convertLinks(ctx context.Context, albumId int) ([]link.RsLink, error) {
	links, err := s.links.LinksFor(ctx, albumId)
	if err != nil {
		return nil, err
	}

	rlinks := make([]link.RsLink, 0)
	for _, ll := range links {
		rlinks = append(rlinks, link.RsLink(ll))
	}
	return rlinks, nil

}

func (s *service) albumsOf(c context.Context, mType MType) ([]RsAlbum, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	albums, err := s.albums.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]RsAlbum, 0)
	for _, album := range albums {
		if album.MType == mType {
			rlinks, err := s.convertLinks(ctx, album.Id)
			if err != nil {
				log.Fatalf("%v", err.Error())
			} else {
				result = append(result, RsAlbum{
					Id:    album.Id,
					Name:  album.Name,
					MType: album.MType,
					Year:  album.Year,
					Slug:  album.Slug,
					Links: rlinks,
				})
			}
		}
	}
	return result, nil
}

func (s *service) OnlyAlbums(ctx context.Context) ([]RsAlbum, error) {
	return s.albumsOf(ctx, AlbumType)
}

func (s *service) OnlySingles(ctx context.Context) ([]RsAlbum, error) {
	return s.albumsOf(ctx, SingleType)
}

func (s *service) AllAlbums(c context.Context) ([]RsAlbum, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	albums, err := s.albums.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]RsAlbum, 0)
	for _, album := range albums {
		rlinks, err := s.convertLinks(ctx, album.Id)
		if err != nil {
			log.Fatalf("%v", err.Error())
		} else {
			result = append(result, RsAlbum{
				Id:    album.Id,
				Name:  album.Name,
				MType: album.MType,
				Year:  album.Year,
				Slug:  album.Slug,
				Links: rlinks,
			})
		}
	}
	return result, nil
}

func (s *service) SongsInAlbum(c context.Context, albumId int) ([]songs.Song, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	songs, err := s.songs.FindAllForAlbum(ctx, albumId)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
