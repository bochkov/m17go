package gigs

import (
	"context"
	"time"

	"github.com/bochkov/m17go/internal/lib/db"
)

type repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) findSince(ctx context.Context, since time.Time) ([]Gig, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, datetime, place_id, description, url
		FROM gigs
		WHERE datetime >= $1
		ORDER BY datetime desc;`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Gig, 0)
	for rows.Next() {
		var g Gig
		rows.Scan(&g.Id, &g.DateTime, &g.Place, &g.Description, &g.Url)
		result = append(result, g)
	}
	return result, nil
}
