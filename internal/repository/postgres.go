package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

type PostgresRepository[Model any] struct {
	db      *sql.DB
	builder *SQLBuilder[Model]
}

func NewPostgresRepository[Model any](db *sql.DB) *PostgresRepository[Model] {
	operators := map[string]func(string, ...any) string{
		// TODO: operators must be refactored based on field type
		// TODO: operators must be synced to schema where types
		"_eq":  func(key string, values ...any) string { return fmt.Sprintf("%s = %s", key, values[0]) },
		"_neq": func(key string, values ...any) string { return fmt.Sprintf("%s != %s", key, values[0]) },
		"_gt":  func(key string, values ...any) string { return fmt.Sprintf("%s > %s", key, values[0]) },
		"_gte": func(key string, values ...any) string { return fmt.Sprintf("%s >= %s", key, values[0]) },
		"_lt":  func(key string, values ...any) string { return fmt.Sprintf("%s < %s", key, values[0]) },
		"_lte": func(key string, values ...any) string { return fmt.Sprintf("%s <= %s", key, values[0]) },
	}
	generator := func(field reflect.StructField, keys *[]any) string {
		return "DEFAULT"
	}
	parameter := func(value reflect.Value, args *[]any) string {
		*args = append(*args, value.Interface())
		return fmt.Sprintf("$%d", len(*args))
	}
	identifier := func(name string) string {
		return fmt.Sprintf("\"%s\"", name)
	}

	return &PostgresRepository[Model]{
		db:      db,
		builder: NewSQLBuilder[Model](operators, generator, parameter, identifier),
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

	// TODO: must use slog with multiple log severities
	fmt.Println(query, args)

	return r.builder.Scan(r.db.QueryContext(ctx, query, args...))
}

func (r *PostgresRepository[Model]) Put(ctx context.Context, models *[]Model) ([]Model, error) {
	// TODO: PK fields must be null
	result := []Model{}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, model := range *models {
		args := []any{}
		query := fmt.Sprintf("UPDATE %s SET %s", r.builder.Table(), r.builder.Set(&model, &args))
		query += fmt.Sprintf(" RETURNING %s", r.builder.Fields())

		fmt.Println(query, args)

		items, err := r.builder.Scan(tx.QueryContext(ctx, query, args...))
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		result = append(result, items...)
	}

	tx.Commit()
	return result, nil
}

func (r *PostgresRepository[Model]) Post(ctx context.Context, models *[]Model) ([]Model, error) {
	args := []any{}
	keys := []any{}
	query := fmt.Sprintf("INSERT INTO %s(%s)", r.builder.Table(), r.builder.Fields())
	if expr := r.builder.Values(models, &keys, &args); expr != "" {
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
