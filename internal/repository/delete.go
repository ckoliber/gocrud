package repository

import (
	"database/sql"
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

func (r *SQLRepository[Model]) Delete(where *map[string]any) ([]Model, error) {
	var rows *sql.Rows
	var err error
	switch r.flavor {
	case sqlbuilder.PostgreSQL, sqlbuilder.SQLite:
		rows, err = r.DeleteReturn(where)
	case sqlbuilder.SQLServer:
		rows, err = r.DeleteOutput(where)
	default:
		rows, err = r.DeleteSelect(where)
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

func (r *SQLRepository[Model]) DeleteReturn(where *map[string]any) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).DeleteFrom(r.table)
	if value := WhereToString(&builder.Cond, where); value != "" {
		builder.Where(value)
	}
	builder.SQL("RETURNING " + strings.Join(r.model.Columns(), ","))

	query, args := builder.Build()

	return r.db.Query(query, args...)
}

func (r *SQLRepository[Model]) DeleteOutput(where *map[string]any) (*sql.Rows, error) {
	builder := r.model.For(r.flavor).DeleteFrom(r.table)
	if value := WhereToString(&builder.Cond, where); value != "" {
		builder.Where(value)
	}
	outputs := []string{}
	for _, column := range r.model.Columns() {
		outputs = append(outputs, "DELETED."+column)
	}
	builder.SQL("OUTPUT " + strings.Join(outputs, ","))

	query, args := builder.Build()

	return r.db.Query(query, args...)
}

func (r *SQLRepository[Model]) DeleteSelect(where *map[string]any) (*sql.Rows, error) {
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
