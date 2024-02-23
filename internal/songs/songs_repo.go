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
		`SELECT s.id, al.pos, s.song_name, s.slug, l.lyrics 
		 FROM songs s, lyrics l, music m, song_on_album al 
		 WHERE m.id = al.music_id and al.song_id = s.id and s.id = l.song_id 
		 AND m.slug = $1
		 ORDER BY al.pos`)
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
		`SELECT s.id, s.song_name, s.slug, soa.pos, l.lyrics, m.name, m.slug, m.year 
		 FROM songs s, lyrics l, music m, song_on_album soa 
		 WHERE s.id = l.song_id
			AND m.id = (
				SELECT soa2.music_id 
				FROM song_on_album soa2, music m2 
				WHERE m2.id = soa2.music_id AND soa2.song_id = s.id AND m2.type = 1 AND m2.ignore = false
				ORDER BY m2.year 
				LIMIT 1
			)
			AND soa.song_id = s.id AND soa.music_id = m.id
		ORDER BY m.year, soa.pos;`)
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
