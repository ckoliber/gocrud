package repository

import (
	"fmt"
	"strings"
)

func (r *CRUDRepository[Model]) Create(models *[]Model) ([]Model, error) {
	builder := r.model.InsertInto(r.table, (*models)[0])
	builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

	query, args := builder.Build()

	fmt.Println(query, args)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Model
	for rows.Next() {
		var model Model
		if err := rows.Scan(r.model.Addr(&model)...); err != nil {
			return nil, err
		}
		result = append(result, model)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
