package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

func ModelsToAnys[Model any](models []Model) []any {
	anySlice := make([]any, len(models))

	for i, model := range models {
		anySlice[i] = model
	}

	return anySlice
}

func (r *SQLRepository[Model]) Post(ctx context.Context, models *[]Model) ([]Model, error) {
	var rows *sql.Rows
	var err error
	switch r.flavor {
	case sqlbuilder.PostgreSQL, sqlbuilder.SQLite:
		rows, err = r.PostReturn(ctx, models)
	case sqlbuilder.SQLServer:
		rows, err = r.PostOutput(ctx, models)
	default:
		rows, err = r.PostSelect(ctx, models)
	}

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

func (r *SQLRepository[Model]) PostReturn(ctx context.Context, models *[]Model) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).WithoutTag("pk").InsertInto(r.table, ModelsToAnys(*models)...)
	builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

	query, args := builder.Build()

	return r.db.QueryContext(ctx, query, args...)
}

func (r *SQLRepository[Model]) PostOutput(ctx context.Context, models *[]Model) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).WithoutTag("pk").InsertInto(r.table, ModelsToAnys(*models)...)

	outputs := []string{}
	for _, column := range r.model.Columns() {
		outputs = append(outputs, "INSERTED."+column)
	}
	builder.SQL("OUTPUT " + strings.Join(outputs, ","))

	query, args := builder.Build()

	return r.db.QueryContext(ctx, query, args...)
}

func (r *SQLRepository[Model]) PostSelect(ctx context.Context, models *[]Model) (*sql.Rows, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	builder := r.model.For(r.flavor).SelectFrom(r.table)
	if value := WhereToString(&builder.Cond, where); value != "" {
		builder.Where(value)
	}
	query, args := builder.Build()
	rows, err := tx.QueryContext(ctx, query, args...)

	deleteBuilder := r.model.For(r.flavor).DeleteFrom(r.table)
	deleteBuilder.WhereClause = builder.WhereClause
	deleteQuery, deleteArgs := deleteBuilder.Build()
	if _, err := tx.ExecContext(ctx, deleteQuery, deleteArgs...); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return rows, err
}
