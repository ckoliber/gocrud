package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgresRepository[Model any] struct {
	db      *sql.DB
	builder *SQLBuilder[Model]
}

func NewPostgresRepository[Model any](db *sql.DB) *PostgresRepository[Model] {
	operators := map[string]func(string, ...any) string{
		"_eq":  func(key string, values ...any) string { return fmt.Sprintf("%s = %s", key, values[0]) },
		"_neq": func(key string, values ...any) string { return fmt.Sprintf("%s != %s", key, values[0]) },
		"_gt":  func(key string, values ...any) string { return fmt.Sprintf("%s > %s", key, values[0]) },
		"_gte": func(key string, values ...any) string { return fmt.Sprintf("%s >= %s", key, values[0]) },
		"_lt":  func(key string, values ...any) string { return fmt.Sprintf("%s < %s", key, values[0]) },
		"_lte": func(key string, values ...any) string { return fmt.Sprintf("%s <= %s", key, values[0]) },
	}
	parameter := func(value any, args *[]any) string {
		// *args = append(*args, value)
		// return "?"
		return fmt.Sprintf("'%s'", value)
	}
	identifier := func(name string) string {
		return fmt.Sprintf("\"%s\"", name)
	}

	return &PostgresRepository[Model]{
		db:      db,
		builder: NewSQLBuilder[Model](operators, parameter, identifier),
	}
}

func (r *PostgresRepository[Model]) Get(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) ([]Model, error) {
	args := []any{}
	query := fmt.Sprintf("SELECT %s FROM %s", r.builder.Fields(), r.builder.Table())
	if expr := r.builder.Where(where, &args); expr != "" {
		query += fmt.Sprintf(" WHERE %s", expr)
	}
	if expr := r.builder.Order(order); expr != "" {
		query += fmt.Sprintf(" ORDER BY %s", expr)
	}
	if limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *limit)
	}
	if skip != nil {
		query += fmt.Sprintf(" OFFSET %d", *skip)
	}

	fmt.Println(query, args)

	return r.builder.Scan(r.db.QueryContext(ctx, query, args...))
}

func (r *PostgresRepository[Model]) Put(ctx context.Context, models *[]Model) ([]Model, error) {
	args := []any{}
	query := "UPDATE " + r.builder.Table() + " SET " + "" + " WHERE " // + r.builder.Where(order) // + " LIMIT " + limit + " OFFSET " + skip
	fmt.Println(query, args)
	return nil, nil
}

func (r *PostgresRepository[Model]) Post(ctx context.Context, models *[]Model) ([]Model, error) {
	args := []any{}
	query := fmt.Sprintf("INSERT INTO %s(%s)", r.builder.Table(), r.builder.Fields())
	if expr := r.builder.Values(models, &args); expr != "" {
		query += fmt.Sprintf(" VALUES %s", expr)
	}
	query += fmt.Sprintf(" RETURNING %s", r.builder.Fields())

	fmt.Println(query, args)

	return r.builder.Scan(r.db.QueryContext(ctx, query, args...))
}

func (r *PostgresRepository[Model]) Delete(ctx context.Context, where *map[string]any) ([]Model, error) {
	args := []any{}
	query := fmt.Sprintf("DELETE FROM %s", r.builder.Table())
	if expr := r.builder.Where(where, &args); expr != "" {
		query += fmt.Sprintf(" WHERE %s", expr)
	}
	query += fmt.Sprintf(" RETURNING %s", r.builder.Fields())

	fmt.Println(query, args)

	return r.builder.Scan(r.db.QueryContext(ctx, query, args...))
}
