package repository

import (
	"database/sql"
	"log/slog"
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

func (r *SQLRepository[Model]) Post(models *[]Model) ([]Model, error) {
	var rows *sql.Rows
	var err error
	switch r.flavor {
	case sqlbuilder.PostgreSQL, sqlbuilder.SQLite:
		rows, err = r.PostReturn(models)
	case sqlbuilder.SQLServer:
		rows, err = r.PostOutput(models)
	default:
		rows, err = r.PostSelect(models)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	slog.Debug()

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

func (r *SQLRepository[Model]) PostReturn(models *[]Model) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).WithoutTag("pk").InsertInto(r.table, ModelsToAnys(*models)...)
	builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

	query, args := builder.Build()

	return r.db.Query(query, args...)
}

func (r *SQLRepository[Model]) PostOutput(models *[]Model) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).WithoutTag("pk").InsertInto(r.table, ModelsToAnys(*models)...)

	outputs := []string{}
	for _, column := range r.model.Columns() {
		outputs = append(outputs, "INSERTED."+column)
	}
	builder.SQL("OUTPUT " + strings.Join(outputs, ","))

	query, args := builder.Build()

	return r.db.Query(query, args...)
}

func (r *SQLRepository[Model]) PostSelect(models *[]Model) (*sql.Rows, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	builder := r.model.For(r.flavor).SelectFrom(r.table)
	if value := WhereToString(&builder.Cond, where); value != "" {
		builder.Where(value)
	}
	query, args := builder.Build()
	rows, err := tx.Query(query, args...)

	deleteBuilder := r.model.For(r.flavor).DeleteFrom(r.table)
	deleteBuilder.WhereClause = builder.WhereClause
	deleteQuery, deleteArgs := deleteBuilder.Build()
	if _, err := tx.Exec(deleteQuery, deleteArgs...); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return rows, err
}
