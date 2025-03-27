package repository

import "fmt"

func (r *CRUDRepository[Model]) Read(where *map[string]any, order *map[string]string, limit *int, skip *int) ([]Model, error) {
	builder := r.model.SelectFrom(r.table)
	if where != nil {
		builder.Where(WhereToString(&builder.Cond, *where))
	}
	if order != nil {
		builder.OrderBy(OrderToString(*order))
	}
	if limit != nil {
		builder.Limit(*limit)
	}
	if skip != nil {
		builder.Offset(*skip)
	}

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
