package repository

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
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

	// TODO: fill models pk fields (Auto generate from type)
	for _, model := range *models {
		_type := reflect.TypeOf(model)
		_value := reflect.ValueOf(model)
		for i := range _type.NumField() {
			field := _type.Field(i)
			value := _value.Field(i)
			if tag := field.Tag.Get(sqlbuilder.FieldTag); strings.Contains(tag, "pk") {
				switch value.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					value.SetInt(1)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					value.SetUint(1)
				case reflect.Float32, reflect.Float64:
					value.SetFloat(1.1)
				case reflect.Complex64, reflect.Complex128:
					value.SetComplex(1)
				case reflect.String:
					value.SetString("1")
				default:
					return nil, errors.New("invalid ID type")
				}
			}
		}
	}

	mutateBuilder := r.model.For(r.flavor).InsertInto(r.table, ModelsToAnys(*models)...)
	mutateQuery, mutateArgs := mutateBuilder.Build()
	if _, err := tx.ExecContext(ctx, mutateQuery, mutateArgs...); err != nil {
		tx.Rollback()
		return nil, err
	}

	builder := r.model.For(r.flavor).SelectFrom(r.table)
	builder.Where(builder.In("id", []string{"1", "2"}))
	query, args := builder.Build()
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return rows, err
}
