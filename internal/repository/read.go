package repository

import (
	"fmt"

	"github.com/ckoliber/gocrud/internal/schema"
)

func (r *CRUDRepository[Model]) Read(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int) ([]Model, error) {
	builder := r.model.SelectFrom(r.table)
	builder.Where(where)

	query, args := builder.Build()

	rows, err := r.db.Query(query, args)
	if err != nil {
		return nil, err
	}

	var result []Model
	rows.Scan(result)

	fmt.Println(query)
	fmt.Println(args)

	return result, nil
}
