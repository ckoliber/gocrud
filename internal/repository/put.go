package repository

import (
	"context"
	"database/sql"
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

func (r *SQLRepository[Model]) Put(ctx context.Context, models *[]Model) ([]Model, error) {
	result := []Model{}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, model := range *models {
		var rows *sql.Rows
		var err error
		switch r.flavor {
		case sqlbuilder.PostgreSQL, sqlbuilder.SQLite:
			rows, err = r.PutReturn(ctx, tx, model)
		case sqlbuilder.SQLServer:
			rows, err = r.PutOutput(ctx, tx, model)
		default:
			rows, err = r.PutSelect(ctx, tx, model)
		}

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

func (r *SQLRepository[Model]) PutReturn(ctx context.Context, tx *sql.Tx, model Model) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).WithoutTag("pk").Update(r.table, model)
	builder.Where(WhereToString(&builder.Cond, ModelToWhere(r.model.WithTag("pk"), model)))
	builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

	query, args := builder.Build()

	return tx.QueryContext(ctx, query, args...)
}

func (r *SQLRepository[Model]) PutOutput(ctx context.Context, tx *sql.Tx, model Model) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).WithoutTag("pk").Update(r.table, model)
	builder.Where(WhereToString(&builder.Cond, ModelToWhere(r.model.WithTag("pk"), model)))

	outputs := []string{}
	for _, column := range r.model.Columns() {
		outputs = append(outputs, "UPDATED."+column)
	}
	builder.SQL("OUTPUT " + strings.Join(outputs, ","))

	query, args := builder.Build()

	return tx.QueryContext(ctx, query, args...)
}

func (r *SQLRepository[Model]) PutSelect(ctx context.Context, tx *sql.Tx, model Model) (*sql.Rows, error) {
	mutateBuilder := r.model.For(r.flavor).WithoutTag("pk").Update(r.table, model)
	mutateBuilder.Where(WhereToString(&mutateBuilder.Cond, ModelToWhere(r.model.WithTag("pk"), model)))
	mutateQuery, mutateArgs := mutateBuilder.Build()
	if _, err := tx.ExecContext(ctx, mutateQuery, mutateArgs...); err != nil {
		return nil, err
	}

	builder := r.model.For(r.flavor).SelectFrom(r.table)
	builder.WhereClause = mutateBuilder.WhereClause
	query, args := builder.Build()
	return tx.QueryContext(ctx, query, args...)
}
