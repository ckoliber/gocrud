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
	identifier := func(name string) string {
		return fmt.Sprintf("\"%s\"", name)
	}
	parameter := func(value any, args *[]any) string {
		*args = append(*args, value)
		return "?"
	}

	return &PostgresRepository[Model]{
		db:      db,
		builder: NewSQLBuilder[Model](operators, parameter, identifier),
	}
}

func (r *PostgresRepository[Model]) Get(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) ([]Model, error) {
	args := []any{}
	query := "SELECT " + r.builder.Columns() + " FROM " + r.builder.table + " WHERE " + r.builder.Where(where, &args) + " ORDER BY " + r.builder.Order(order) // + " LIMIT " + limit + " OFFSET " + skip
	fmt.Println(query, args)
	return nil, nil
}

func (r *PostgresRepository[Model]) Put(ctx context.Context, models *[]Model) ([]Model, error) {
	args := []any{}
	query := "UPDATE " + r.builder.table + " SET " + "" + " WHERE " // + r.builder.Where(order) // + " LIMIT " + limit + " OFFSET " + skip
	fmt.Println(query, args)
	return nil, nil
}

func (r *PostgresRepository[Model]) Post(ctx context.Context, models *[]Model) ([]Model, error) {
	args := []any{}
	query := "INSERT INTO " + r.builder.table + "(" + r.builder.Columns() + ")" + " VALUES " + r.builder.Values(models, &args)
	fmt.Println(query, args)
	return nil, nil
}

func (r *PostgresRepository[Model]) Delete(ctx context.Context, where *map[string]any) ([]Model, error) {
	args := []any{}
	query := "DELETE FROM " + r.builder.table + " WHERE " + r.builder.Where(where, &args)
	fmt.Println(query, args)
	return nil, nil
}
