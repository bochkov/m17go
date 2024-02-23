package albums

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

func (r *repository) FindLatest(ctx context.Context) (*Album, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT mus.id, mus.name, mus.year, mus.type, mus.slug 
		 FROM music mus 
		 WHERE mus.id = (SELECT max(mu.id) FROM music mu WHERE mu.ignore=false)`)

	var a Album
	if err := row.Scan(&a.Id, &a.Name, &a.Year, &a.MType, &a.Slug); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *repository) FindAll(ctx context.Context) ([]Album, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT mus.id, mus.name, mus.year, mus.type, mus.slug
		 FROM music mus
		 WHERE mus.ignore=false
		 ORDER by mus.year desc, mus.id desc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allAlbums := make([]Album, 0)
	for rows.Next() {
		var a Album
		rows.Scan(&a.Id, &a.Name, &a.Year, &a.MType, &a.Slug)
		allAlbums = append(allAlbums, a)
	}
	return allAlbums, nil
}

func (r *repository) FindBySlug(ctx context.Context, albumSlug string) (*Album, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT mus.id, mus.name, mus.year, mus.type, mus.slug
		 FROM music mus
		 WHERE mus.slug=$1`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, albumSlug)

	var a Album
	row.Scan(&a.Id, &a.Name, &a.Year, &a.MType, &a.Slug)
	return &a, nil
}
