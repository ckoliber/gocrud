package repository

import (
	"database/sql"
	"reflect"
	"strings"
)

type CRUDRepository[Model any] struct {
	db    *sql.DB
	table string
}

func NewCRUDRepository[Model any](db *sql.DB) *CRUDRepository[Model] {
	_type := reflect.TypeFor[Model]()

	result := &CRUDRepository[Model]{
		db:    db,
		table: strings.ToLower(_type.Name()),
		// model: sqlbuilder.NewStruct(new(Model)),
	}

	if field, ok := _type.FieldByName("_"); ok {
		if value := field.Tag.Get("table"); value != "" {
			result.table = value
		}
	}

	return result
}

func (r *CRUDRepository[Model]) Transaction() (*sql.Tx, error) {
	return r.db.Begin()
}
