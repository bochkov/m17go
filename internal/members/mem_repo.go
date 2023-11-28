package members

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

func (r *repository) findAll(ctx context.Context) ([]Member, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT me.id, me.name, inst.text as instrument, me.actual 
		 FROM members me, instrument inst 
		 WHERE me.instrument = inst.id 
		 ORDER BY me.weight`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Member, 0)
	for rows.Next() {
		var m Member
		rows.Scan(&m.Id, &m.Name, &m.Instrument, &m.Actual)
		result = append(result, m)
	}
	return result, nil
}

func (r *repository) findActual(ctx context.Context) ([]Member, error) {
	var result []Member
	members, err := r.findAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, mem := range members {
		if mem.Actual {
			result = append(result, mem)
		}
	}
	return result, nil
}
