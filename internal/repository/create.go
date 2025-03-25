package repository

import (
	"fmt"

	"github.com/ckoliber/gocrud/internal/schema"
)

func (r *CRUDRepository[Model]) Create(fields *schema.Fields[Model], models *[]Model) ([]Model, error) {
	builder := r.model.InsertInto(r.table, models).Returning(r.model.Columns()...)

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
