package repository

import (
	"database/sql"
	"reflect"
	"strings"
	"text/template"
)

type CRUDRepository[Model any] struct {
	db       *sql.DB
	table    string
	columns  []string
	template *template.Template
}

func NewCRUDRepository[Model any](db *sql.DB, SQL *template.Template) *CRUDRepository[Model] {
	_type := reflect.TypeFor[Model]()

	result := &CRUDRepository[Model]{
		db:       db,
		table:    strings.ToLower(_type.Name()),
		columns:  []string{},
		template: SQL,
	}

	if SQL == nil {
		switch reflect.TypeOf(db.Driver()).String() {
		case "*pq.Driver", "*pqx.Driver":
			result.template = getPostgres()
		case "*mysql.MySQLDriver":
			result.template = getMySQL()
		case "*mssql.MssqlDriver":
			result.template = getMSSQL()
		case "*sqlite.SQLiteDriver":
			result.template = getSQLite()
		}
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
