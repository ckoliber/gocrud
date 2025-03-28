package repository

func (r *CRUDRepository[Model]) Read(where *map[string]any, order *map[string]string, limit *int, skip *int) ([]Model, error) {
	builder := r.model.SelectFrom(r.table)
	if value := WhereToString(&builder.Cond, where); value != "" {
		builder.Where(value)
	}
	if value := OrderToString(order); value != "" {
		builder.OrderBy(value)
	}
	if limit != nil {
		builder.Limit(*limit)
	}
	if skip != nil {
		builder.Offset(*skip)
	}

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
