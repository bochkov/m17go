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

func (r *repository) FindAllForAlbum(ctx context.Context, albumId int) ([]Song, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT s.id, al.pos, s.song_name, l.lyrics 
		 FROM songs s, lyrics l, music m, song_on_album al 
		 WHERE m.id = al.music_id and al.song_id = s.id and s.id = l.song_id 
		 AND m.id = $1
		 ORDER BY al.pos`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, albumId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	songs := make([]Song, 0)
	for rows.Next() {
		var s Song
		rows.Scan(&s.Id, &s.Position, &s.Name, &s.Lyrics)
		songs = append(songs, s)
	}

	return songs, nil
}
