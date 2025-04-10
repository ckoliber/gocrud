package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
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

type Relation struct {
	one   bool
	src   string
	dest  string
	table string
}

type SQLBuilder[Model any] struct {
	table      string
	keys       []string
	fields     []Field
	relations  map[string]Relation
	operations map[string]func(string, ...string) string
	identifier func(string) string
	parameter  func(reflect.Value, *[]any) string
	generator  func(reflect.StructField, *[]any) string
}

type SQLBuilderInterface interface {
	Table() string
	Where(where *map[string]any, args *[]any, run func(string) []string) string
}

var registry = map[string]SQLBuilderInterface{}

func NewSQLBuilder[Model any](operations map[string]func(string, ...string) string, identifier func(string) string, parameter func(reflect.Value, *[]any) string, generator func(reflect.StructField, *[]any) string) *SQLBuilder[Model] {
	// Initialize SQLBuilder with table name and fields based on the Model type
	// Logs the table name and fields for debugging
	_type := reflect.TypeFor[Model]()

	table := strings.ToLower(_type.Name())
	fields := []Field{}
	relations := map[string]Relation{}
	for idx := range _type.NumField() {
		_field := _type.Field(idx)
		if _field.Name == "_" {
			if tag := _field.Tag.Get("db"); tag != "" {
				table = strings.Split(tag, ",")[0]
			}
		} else {
			if tag := _field.Tag.Get("db"); tag != "" {
				if _field.Tag.Get("json") == "-" {
					relations[tag] = Relation{
						one:   _field.Type.Kind() == reflect.Struct,
						src:   _field.Tag.Get("src"),
						dest:  _field.Tag.Get("dest"),
						table: _field.Tag.Get("table"),
					}
				} else {
					fields = append(fields, Field{idx, strings.Split(tag, ",")[0]})
				}
			}
		}
	}

	slog.Debug("SQLBuilder initialized", slog.String("table", table), slog.Any("fields", fields), slog.Any("relations", relations))

	result := &SQLBuilder[Model]{
		table:      table,
		keys:       []string{fields[0].name},
		fields:     fields,
		relations:  relations,
		operations: operations,
		identifier: identifier,
		parameter:  parameter,
		generator:  generator,
	}

	registry[table] = result

	return result
}

func (b *SQLBuilder[Model]) Table() string {
	// Returns the table name with proper identifier formatting
	slog.Debug("Fetching table name", slog.String("table", b.table))
	return b.identifier(b.table)
}

func (b *SQLBuilder[Model]) Fields(prefix string) string {
	// Returns a comma-separated list of field names with proper identifier formatting
	result := []string{}
	for _, field := range b.fields {
		result = append(result, prefix+b.identifier(field.name))
	}
	slog.Debug("Fetching fields", slog.Any("fields", result))
	return strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Values(values *[]Model, args *[]any, keys *[]any) (string, string) {
	// Constructs the VALUES clause for an INSERT query
	if values == nil {
		return "", ""
	}

	fields := []string{}
	for idx, field := range b.fields {
		if idx == 0 {
			if b.generator != nil {
				fields = append(fields, b.identifier(field.name))
			}
		} else {
			fields = append(fields, b.identifier(field.name))
		}
	}

	result := []string{}
	for _, model := range *values {
		_type := reflect.TypeOf(model)
		_value := reflect.ValueOf(model)

		items := []string{}
		for idx, field := range b.fields {
			if idx == 0 {
				if b.generator != nil {
					items = append(items, b.generator(_type.Field(field.idx), keys))
				}
			} else {
				items = append(items, b.parameter(_value.Field(field.idx), args))
			}
		}

		result = append(result, "("+strings.Join(items, ",")+")")
	}

	slog.Debug("Constructed VALUES clause", slog.Any("values", result))
	return strings.Join(fields, ","), strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Set(set *Model, args *[]any, where *map[string]any) string {
	// Constructs the SET clause for an UPDATE query
	if set == nil {
		return ""
	}

	_value := reflect.ValueOf(*set)

	result := []string{}
	for idx, field := range b.fields {
		if idx == 0 {
			if where != nil {
				_field := _value.Field(field.idx)
				for _field.Kind() == reflect.Pointer {
					_field = _field.Elem()
				}

				switch _field.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					(*where)[field.name] = map[string]any{"_eq": fmt.Sprintf("%d", _field.Int())}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					(*where)[field.name] = map[string]any{"_eq": fmt.Sprintf("%d", _field.Uint())}
				case reflect.Float32, reflect.Float64:
					(*where)[field.name] = map[string]any{"_eq": fmt.Sprintf("%f", _field.Float())}
				case reflect.Complex64, reflect.Complex128:
					(*where)[field.name] = map[string]any{"_eq": fmt.Sprintf("%f", _field.Complex())}
				case reflect.String:
					(*where)[field.name] = map[string]any{"_eq": _field.String()}
				default:
					panic("Invalid identifier type")
				}
			}
		} else {
			result = append(result, field.name+"="+b.parameter(_value.Field(field.idx), args))
		}
	}

	slog.Debug("Constructed SET clause", slog.String("set", strings.Join(result, ",")))
	return strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Order(order *map[string]any) string {
	// Constructs the ORDER BY clause for a query
	if order == nil {
		return ""
	}

	result := []string{}
	for key, val := range *order {
		result = append(result, fmt.Sprintf("%s %s", b.identifier(key), val))
	}

	slog.Debug("Constructed ORDER BY clause", slog.Any("order", result))
	return strings.Join(result, ",")
}

func (b *SQLBuilder[Model]) Where(where *map[string]any, args *[]any, run func(string) []string) string {
	// Constructs the WHERE clause for a query
	if where == nil {
		return ""
	}

	if item, ok := (*where)["_not"]; ok {
		expr := item.(map[string]any)

		return "NOT (" + b.Where(&expr, args, run) + ")"
	} else if items, ok := (*where)["_and"]; ok {
		result := []string{}
		for _, item := range items.([]any) {
			expr := item.(map[string]any)
			result = append(result, b.Where(&expr, args, run))
		}

		return "(" + strings.Join(result, " AND ") + ")"
	} else if items, ok := (*where)["_or"]; ok {
		result := []string{}
		for _, item := range items.([]any) {
			expr := item.(map[string]any)
			result = append(result, b.Where(&expr, args, run))
		}

		return "(" + strings.Join(result, " OR ") + ")"
	}

	result := []string{}
	for key, item := range *where {
		for op, value := range item.(map[string]any) {
			if handler, ok := b.operations[op]; ok {
				_value := reflect.ValueOf(value)

				if _value.Kind() == reflect.String {
					result = append(result, handler(b.identifier(key), b.parameter(_value, args)))
				} else if _value.Kind() == reflect.Slice || _value.Kind() == reflect.Array {
					items := []string{}
					for i := range _value.Len() {
						items = append(items, b.parameter(_value.Index(i), args))
					}

					result = append(result, handler(b.identifier(key), items...))
				}
			} else {
				if relation, ok := b.relations[key]; ok {
					builder := registry[relation.table]

					args_ := []any{}
					where := item.(map[string]any)
					query := fmt.Sprintf("SELECT %s FROM %s", b.identifier(relation.dest), builder.Table())
					if expr := builder.Where(&where, &args_, run); expr != "" {
						query += fmt.Sprintf(" WHERE %s", expr)
					}

					if run == nil {
						*args = append(*args, args_...)
						result = append(result, b.operations["_in"](b.identifier(relation.src), query))
					} else {
						result = append(result, b.operations["_in"](b.identifier(relation.src), run(query)...))
					}
				}
			}
		}
	}

	slog.Debug("Constructed WHERE clause", slog.Any("where", result))
	return strings.Join(result, " AND ")
}

func (b *SQLBuilder[Model]) Scan(rows *sql.Rows, err error) ([]Model, error) {
	// Scans the rows returned by a query into a slice of Model
	if err != nil {
		slog.Error("Error during query execution", slog.Any("error", err))
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
		slog.Error("Error during row iteration", slog.Any("error", err))
		return nil, err
	}

	slog.Debug("Scan completed", slog.Any("result", result))
	return result, nil
}
