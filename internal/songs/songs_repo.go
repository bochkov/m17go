package songs

import (
	"context"

	"github.com/bochkov/m17go/internal/lib/db"
)

type repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) SongsForAlbum(ctx context.Context, albumSlug string) ([]Song, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT s.id, soa.position, s.name, s.slug, l.text
		FROM songs s, lyrics l, albums a, songs_on_albums soa
		WHERE a.id = soa.album_id and soa.song_id = s.id and s.id = l.song_id
		AND a.slug = $1
		ORDER BY soa.position`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, albumSlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	songs := make([]Song, 0)
	for rows.Next() {
		var s Song
		rows.Scan(&s.Id, &s.Position, &s.Name, &s.Slug, &s.Lyrics)
		songs = append(songs, s)
	}

	return songs, nil
}

func (r *repository) AllSongs(ctx context.Context) ([]SongWithAlbum, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT s.id, s.name, s.slug, soa.position, l.text, a.name, a.slug, a.year
		FROM songs s, lyrics l, albums a, songs_on_albums soa
		WHERE s.id = l.song_id
			AND a.id = (
				SELECT soa2.album_id
				FROM songs_on_albums soa2, albums a2
				WHERE a2.id = soa2.album_id AND soa2.song_id = s.id AND a2.type = 'album' AND a2.ignore = false
				ORDER BY a2.year
				LIMIT 1
			)
			AND soa.song_id = s.id AND soa.album_id = a.id
		ORDER BY a.year, soa.position`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allSongs := make([]SongWithAlbum, 0)
	for rows.Next() {
		var a SongWithAlbum
		rows.Scan(&a.Id, &a.Name, &a.Slug, &a.Position,
			&a.Lyrics, &a.AlbumName, &a.AlbumSlug, &a.AlbumYear)
		allSongs = append(allSongs, a)
	}
	return allSongs, nil
}
