package repository

import (
	"context"
	"database/sql"
	"reflect"
	"strings"
)

type MySQLRepository[Model any] struct {
	db    *sql.DB
	table string
}

func NewMySQLRepository[Model any](db *sql.DB) *MySQLRepository[Model] {
	_type := reflect.TypeFor[Model]()

	result := &MySQLRepository[Model]{
		db:    db,
		table: strings.ToLower(_type.Name()),
		// model:  sqlbuilder.NewStruct(new(Model)),
		// flavor: sqlbuilder.DefaultFlavor,
	}

	if field, ok := _type.FieldByName("_"); ok {
		if value := field.Tag.Get("db"); value != "" {
			result.table = value
		}
	}

	return result
}

func (r *MySQLRepository[Model]) Get(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) ([]Model, error) {
	return nil, nil
}

func (r *MySQLRepository[Model]) Put(ctx context.Context, models *[]Model) ([]Model, error) {
	return nil, nil
}

func (r *MySQLRepository[Model]) Post(ctx context.Context, models *[]Model) ([]Model, error) {
	return nil, nil
}

func (r *MySQLRepository[Model]) Delete(ctx context.Context, where *map[string]any) ([]Model, error) {
	return nil, nil
}
