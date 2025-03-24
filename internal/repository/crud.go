package repository

import "database/sql"

type CRUDRepository[Model any] struct {
	db      *sql.DB
	table   string
	columns []string
}
