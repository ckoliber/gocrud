package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type Repository[Model any] interface {
	Get(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) ([]Model, error)
	Put(ctx context.Context, models *[]Model) ([]Model, error)
	Post(ctx context.Context, models *[]Model) ([]Model, error)
	Delete(ctx context.Context, where *map[string]any) ([]Model, error)
}

type Field struct {
	idx  int
	name string
}

type SQLBuilder[Model any] struct {
	table      string
	fields     []Field
	operators  map[string]func(string, ...any) string
	generator  func(reflect.StructField, *[]any) string
	parameter  func(reflect.Value, *[]any) string
	identifier func(string) string
}

func NewSQLBuilder[Model any](operators map[string]func(string, ...any) string, generator func(reflect.StructField, *[]any) string, parameter func(reflect.Value, *[]any) string, identifier func(string) string) *SQLBuilder[Model] {
	_type := reflect.TypeFor[Model]()

	table := strings.ToLower(_type.Name())
	fields := []Field{}
	for idx := range _type.NumField() {
		if _type.Field(idx).Name == "_" {
			if tag := _type.Field(idx).Tag.Get("db"); tag != "" {
				table = strings.Split(tag, ",")[0]
			}
		} else {
			if tag := _type.Field(idx).Tag.Get("db"); tag != "" {
				fields = append(fields, Field{idx, strings.Split(tag, ",")[0]})
			}
		}
	}

	return &SQLBuilder[Model]{
		table:      table,
		fields:     fields,
		operators:  operators,
		generator:  generator,
		parameter:  parameter,
		identifier: identifier,
	}
}

func (b *SQLBuilder[Model]) Table() string {
	return b.identifier(b.table)
}

func (b *SQLBuilder[Model]) Fields() string {
	result := []string{}
	for _, field := range b.fields {
		result = append(result, b.identifier(field.name))
	}

	return strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Values(values *[]Model, keys *[]any, args *[]any) string {
	if values == nil {
		return ""
	}

	result := []string{}
	for _, model := range *values {
		_type := reflect.TypeOf(model)
		_value := reflect.ValueOf(model)

		fields := []string{}
		for idx, field := range b.fields {
			if idx == 0 {
				fields = append(fields, b.generator(_type.Field(field.idx), keys))
			} else {
				fields = append(fields, b.parameter(_value.Field(field.idx), args))
			}
		}

		result = append(result, "("+strings.Join(fields, ",")+")")
	}

	return strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Set(set *Model, args *[]any) string {
	if set == nil {
		return ""
	}

	_value := reflect.ValueOf(*set)

	result := []string{}
	for _, field := range b.fields {
		result = append(result, field.name+"="+b.parameter(_value.Field(field.idx), args))
	}

	where := map[string]any{}
	for idx, field := range b.fields {
		if idx == 0 {
			where[field.name] = map[string]any{"_eq": _value.Field(field.idx).Interface()}
		}
	}

	return strings.Join(result, ",") + " WHERE " + b.Where(&where, args)
}

func (b *SQLBuilder[Model]) Order(order *map[string]any) string {
	if order == nil {
		return ""
	}

	result := []string{}
	for key, val := range *order {
		result = append(result, fmt.Sprintf("%s %s", b.identifier(key), val))
	}

	return strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Where(where *map[string]any, args *[]any) string {
	if where == nil {
		return ""
	}

	if item, ok := (*where)["_not"]; ok {
		expr := item.(map[string]any)

		return "NOT (" + b.Where(&expr, args) + ")"
	} else if items, ok := (*where)["_and"]; ok {
		result := []string{}
		for _, item := range items.([]any) {
			expr := item.(map[string]any)
			result = append(result, b.Where(&expr, args))
		}

		return "(" + strings.Join(result, " AND ") + ")"
	} else if items, ok := (*where)["_or"]; ok {
		result := []string{}
		for _, item := range items.([]any) {
			expr := item.(map[string]any)
			result = append(result, b.Where(&expr, args))
		}

		return "(" + strings.Join(result, " OR ") + ")"
	}

	result := []string{}
	for key, item := range *where {
		for op, value := range item.(map[string]any) {
			if handler, ok := b.operators[op]; ok {
				result = append(result, handler(b.identifier(key), b.parameter(reflect.ValueOf(value), args)))
			}
		}
	}

	return strings.Join(result, " AND ")
}

func (b *SQLBuilder[Model]) Scan(rows *sql.Rows, err error) ([]Model, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Model{}
	for rows.Next() {
		var model Model
		_value := reflect.ValueOf(&model).Elem()

		_addrs := []any{}
		for _, field := range b.fields {
			_addrs = append(_addrs, _value.Field(field.idx).Addr().Interface())
		}

		if err := rows.Scan(_addrs...); err != nil {
			return nil, err
		}

		result = append(result, model)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
