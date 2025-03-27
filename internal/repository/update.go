package repository

import "strings"

func (r *CRUDRepository[Model]) Update(where *map[string]any, model *Model) ([]Model, error) {
	builder := r.model.Update(r.table, model)
	builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))
	if where != nil {
		builder.Where(WhereToString(&builder.Cond, *where))
	}

	query, args := builder.Build()

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
