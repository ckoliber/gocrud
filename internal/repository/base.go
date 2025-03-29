package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

type Repository[Model any] interface {
	Get(where *map[string]any, order *map[string]any, limit *int, skip *int) ([]Model, error)
	Put(models *[]Model) ([]Model, error)
	Post(models *[]Model) ([]Model, error)
	Delete(where *map[string]any) ([]Model, error)
}

type SQLRepository[Model any] struct {
	db     *sql.DB
	table  string
	model  *sqlbuilder.Struct
	flavor sqlbuilder.Flavor
}

func NewSQLRepository[Model any](db *sql.DB) *SQLRepository[Model] {
	_type := reflect.TypeFor[Model]()

	result := &SQLRepository[Model]{
		db:     db,
		table:  strings.ToLower(_type.Name()),
		model:  sqlbuilder.NewStruct(new(Model)),
		flavor: sqlbuilder.DefaultFlavor,
	}

	if field, ok := _type.FieldByName("_"); ok {
		if value := field.Tag.Get("db"); value != "" {
			result.table = value
		}
	}

	switch reflect.ValueOf(db.Driver()).Type().String() {
	case "*mysql.MySQLDriver":
		result.flavor = sqlbuilder.MySQL
	case "*pq.Driver", "pqx.Driver":
		result.flavor = sqlbuilder.PostgreSQL
	case "*sqlite.SQLiteDriver":
		result.flavor = sqlbuilder.SQLite
	case "*mssql.MssqlDriver":
		result.flavor = sqlbuilder.SQLServer
	}

	return result
}

func OrderToString(order *map[string]any) string {
	if order == nil {
		return ""
	}

	result := []string{}

	for key, val := range *order {
		result = append(result, fmt.Sprintf("%s %s", key, val))
	}

	return strings.Join(result, ",")
}

func WhereToString(cond *sqlbuilder.Cond, where *map[string]any) string {
	if where == nil {
		return ""
	}

	if item, ok := (*where)["_not"]; ok {
		expr := item.(map[string]any)

		return cond.Not(WhereToString(cond, &expr))
	} else if items, ok := (*where)["_and"]; ok {
		result := []string{}
		for _, item := range items.([]any) {
			expr := item.(map[string]any)
			result = append(result, WhereToString(cond, &expr))
		}

		return cond.And(result...)
	} else if items, ok := (*where)["_or"]; ok {
		result := []string{}
		for _, item := range items.([]any) {
			expr := item.(map[string]any)
			result = append(result, WhereToString(cond, &expr))
		}

		return cond.Or(result...)
	}

	result := []string{}
	for key, item := range *where {
		expr := item.(map[string]any)

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
