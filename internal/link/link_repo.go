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

func (r *repository) LinksFor(ctx context.Context, id int) ([]Link, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT ml.id, mp.id, mp.name, ml.url 
		 FROM music_links ml, music_provs mp
		 WHERE ml.provider = mp.id and ml.music = $1
		 ORDER BY mp.id, ml.id`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]Link, 0)
	for rows.Next() {
		var l Link
		err := rows.Scan(&l.Id, &l.ProvId, &l.Provider, &l.Url)
		if err != nil {
			return nil, err
		}
		links = append(links, l)
	}
	return links, nil
}
