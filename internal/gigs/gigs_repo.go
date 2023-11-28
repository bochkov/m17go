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
		`SELECT g.id, g.dt, g.tm, g.place, g.desc, g.url
		FROM gigs g 
		WHERE g.dt >= $1 
		ORDER BY g.dt desc, g.tm`)
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
		rows.Scan(&g.Id, &g.Date, &g.Time, &g.Place, &g.Desc, &g.Url)
		result = append(result, g)
	}
	return result, nil
}
