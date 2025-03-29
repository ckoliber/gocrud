package repository

import (
	"strings"
)

func (r *SQLRepository[Model]) Post(models *[]Model) ([]Model, error) {
	anySlice := make([]any, len(*models))
	for i, model := range *models {
		anySlice[i] = model
	}

	builder := r.model.For(r.flavor).WithoutTag("pk").InsertInto(r.table, anySlice...)
	builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

	query, args := builder.Build()

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Model{}
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
