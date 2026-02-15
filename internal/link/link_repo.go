package link

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

func (r *repository) LinksFor(ctx context.Context, albumId int) ([]Link, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, url
		FROM album_links
		WHERE album_id = $1
		ORDER BY id`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, albumId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]Link, 0)
	for rows.Next() {
		var l Link
		err := rows.Scan(&l.Id, &l.Url)
		if err != nil {
			return nil, err
		}
		l.ProviderId = detectProvider(l.Url)
		links = append(links, l)
	}
	return links, nil
}
