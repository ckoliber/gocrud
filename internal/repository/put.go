package repository

func (r *SQLRepository[Model]) Put(models *[]Model) ([]Model, error) {
	return nil, nil
	// builder := r.model.WithoutTag("pk").Update(r.table, model)
	// if value := WhereToString(&builder.Cond, where); value != "" {
	// 	builder.Where(value)
	// }
	// builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

	// query, args := builder.Build()

	// rows, err := r.db.Query(query, args...)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// result := []Model{}
	// for rows.Next() {
	// 	var model Model
	// 	if err := rows.Scan(r.model.Addr(&model)...); err != nil {
	// 		return nil, err
	// 	}
	// 	result = append(result, model)
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }

	// return result, nil
}
