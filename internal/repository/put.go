package repository

import (
	"fmt"
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

func ModelToWhere[Model any](_struct *sqlbuilder.Struct, model Model) *map[string]any {
	result := map[string]any{}

	columns := _struct.Columns()
	values := _struct.Values(&model)
	for idx := range len(values) {
		result[columns[idx]] = map[string]any{"_eq": values[idx]}
	}

	return &result
}

func (r *SQLRepository[Model]) Put(models *[]Model) ([]Model, error) {
	result := []Model{}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	for _, model := range *models {
		fmt.Println(r.model.Values(&model)...)
		builder := r.model.WithoutTag("pk").Update(r.table, model)
		builder.Where(WhereToString(&builder.Cond, ModelToWhere(r.model.WithTag("pk"), model)))
		builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

		query, args := builder.Build()

		fmt.Println(query, args)

		rows, err := tx.Query(query, args...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var model Model
			if err := rows.Scan(r.model.Addr(&model)...); err != nil {
				tx.Rollback()
				return nil, err
			}
			result = append(result, model)
		}
		if err = rows.Err(); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return result, nil
}
