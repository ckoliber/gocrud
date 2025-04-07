package repository

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"text/template"
)

type MSSQLRepository[Model any] struct {
	db      *sql.DB
	builder *SQLBuilder[Model]
}

func NewMSSQLRepository[Model any](db *sql.DB) *MSSQLRepository[Model] {
	operators := map[string]func(string, ...string) string{
		"_eq":  func(key string, values ...string) string { return fmt.Sprintf("%s = %s", key, values[0]) },
		"_neq": func(key string, values ...string) string { return fmt.Sprintf("%s != %s", key, values[0]) },
		"_gt":  func(key string, values ...string) string { return fmt.Sprintf("%s > %s", key, values[0]) },
		"_gte": func(key string, values ...string) string { return fmt.Sprintf("%s >= %s", key, values[0]) },
		"_lt":  func(key string, values ...string) string { return fmt.Sprintf("%s < %s", key, values[0]) },
		"_lte": func(key string, values ...string) string { return fmt.Sprintf("%s <= %s", key, values[0]) },
	}
	generator := func(field reflect.StructField, keys *[]any) string {
		return "NULL"
	}
	parameter := func(value reflect.Value, args *[]any) string {
		*args = append(*args, value)
		return fmt.Sprintf("@p%d", len(*args))
	}
	identifier := func(name string) string {
		return fmt.Sprintf("[%s]", name)
	}

	return &MSSQLRepository[Model]{
		db:      db,
		builder: NewSQLBuilder[Model](operators, generator, parameter, identifier),
	}
}

func (r *MSSQLRepository[Model]) Get(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) ([]Model, error) {
	return nil, nil
}

func (r *MSSQLRepository[Model]) Put(ctx context.Context, models *[]Model) ([]Model, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result := []Model{}
	for _, model := range *models {
		tpl := template.Must(template.New("insert").Parse(`
			UPDATE {{ table }} SET {{ set }}
			WHERE {{ where }}
			RETURNING {{ columns }}
		`))

		args := []any{}
		var query bytes.Buffer
		err := tpl.Execute(&query, map[string]any{
			"set": r.builder.Set(&model, &args),
			// TODO: "where":   r.builder.Where(where, &args),
			"columns": r.builder.Fields(),
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		models, err := r.builder.Scan(r.db.QueryContext(ctx, query.String(), args...))
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		result = append(result, models...)
	}

	tx.Commit()
	return result, nil
}

func (r *MSSQLRepository[Model]) Post(ctx context.Context, models *[]Model) ([]Model, error) {
	return nil, nil
}

func (r *MSSQLRepository[Model]) Delete(ctx context.Context, where *map[string]any) ([]Model, error) {
	return nil, nil
}
