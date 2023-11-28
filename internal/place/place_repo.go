package place

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

func (r *repository) FindById(ctx context.Context, id int) (*Place, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM places WHERE id = $1")
	if err != nil {
		return &Place{}, err
	}

	var p Place
	row := stmt.QueryRowContext(ctx, id)
	if err := row.Scan(&p.Id, &p.Name, &p.Address, &p.Link, &p.Slug, &p.InvertedLogo); err != nil {
		return &Place{}, err
	}

	return &p, nil
}
