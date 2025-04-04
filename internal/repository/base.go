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

type SQLBuilder[Model any] struct {
	table      string
	fields     map[int]string
	operators  map[string]func(string, ...any) string
	parameter  func(any, *[]any) string
	identifier func(string) string
}

func NewSQLBuilder[Model any](operators map[string]func(string, ...any) string, parameter func(any, *[]any) string, identifier func(string) string) *SQLBuilder[Model] {
	_type := reflect.TypeFor[Model]()

	fields := map[int]string{}
	for i := range _type.NumField() {
		if _type.Field(i).Name == "_" {
			continue
		}
		if tag := _type.Field(i).Tag.Get("db"); tag != "" {
			fields[i] = tag
		} else if tag := _type.Field(i).Tag.Get("json"); tag != "" {
			fields[i] = tag
		}
	}

	return &SQLBuilder[Model]{
		table:      strings.ToLower(_type.Name()),
		fields:     fields,
		operators:  operators,
		parameter:  parameter,
		identifier: identifier,
	}
}

func (b *SQLBuilder[Model]) Columns() string {
	result := []string{}
	for _, name := range b.fields {
		result = append(result, b.identifier(name))
	}

	return strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Values(values *[]Model, args *[]any) string {
	if values == nil {
		return ""
	}

	result := []string{}
	for _, model := range *values {
		_value := reflect.ValueOf(model)

		fields := []string{}
		for idx := range b.fields {
			fields = append(fields, b.parameter(_value.Field(idx).Interface(), args))
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
	for idx, name := range b.fields {
		result = append(result, name+" = "+b.parameter(_value.Field(idx).Interface(), args))
	}

	return strings.Join(result, ",")
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
		_value := reflect.ValueOf(model)

		_addrs := []any{}
		for idx := range b.fields {
			_addrs = append(_addrs, _value.Field(idx).Addr())
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
