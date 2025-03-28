package repository

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

type Repository[Model any] interface {
	Get(where *map[string]any, order *map[string]string, limit *int, skip *int) ([]Model, error)
	Put(models *[]Model) ([]Model, error)
	Post(models *[]Model) ([]Model, error)
	Delete(where *map[string]any) ([]Model, error)
}

type SQLRepository[Model any] struct {
	db    *sql.DB
	table string
	model *sqlbuilder.Struct
}

func NewSQLRepository[Model any](db *sql.DB) *SQLRepository[Model] {
	_type := reflect.TypeFor[Model]()

	result := &SQLRepository[Model]{
		db:    db,
		table: strings.ToLower(_type.Name()),
		model: sqlbuilder.NewStruct(new(Model)).For(sqlbuilder.PostgreSQL),
	}

	if field, ok := _type.FieldByName("_"); ok {
		if value := field.Tag.Get("db"); value != "" {
			result.table = value
		}
	}

	return result
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func OrderToString(order *map[string]string) string {
	if order == nil {
		return ""
	}

	result := []string{}

	for key, val := range *order {
		result = append(result, key+" "+val)
	}

	return strings.Join(result, ",")
}

func WhereToString(cond *sqlbuilder.Cond, where *map[string]any) string {
	if where == nil {
		return ""
	}

	if item, ok := (*where)["_not"]; ok {
		return "NOT (" + WhereToString(cond, item.(*map[string]any)) + ")"
	} else if items, ok := (*where)["_and"]; ok {
		return "(" + strings.Join(Map(items.([]*map[string]any), func(item *map[string]any) string { return WhereToString(cond, item) }), " AND ") + ")"
	} else if items, ok := (*where)["_or"]; ok {
		return "(" + strings.Join(Map(items.([]*map[string]any), func(item *map[string]any) string { return WhereToString(cond, item) }), " OR ") + ")"
	}

	result := []string{}
	for key, val := range *where {
		expr := val.(map[string]any)

		if value, ok := expr["_eq"]; ok {
			result = append(result, cond.EQ(key, value))
		} else if value, ok := expr["_neq"]; ok {
			result = append(result, cond.NEQ(key, value))
		} else if value, ok := expr["_gt"]; ok {
			result = append(result, cond.GT(key, value))
		} else if value, ok := expr["_gte"]; ok {
			result = append(result, cond.GTE(key, value))
		} else if value, ok := expr["_lt"]; ok {
			result = append(result, cond.LT(key, value))
		} else if value, ok := expr["_lte"]; ok {
			result = append(result, cond.LTE(key, value))
		}
	}

	return cond.And(result...)
}
