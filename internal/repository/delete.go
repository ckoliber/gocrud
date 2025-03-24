package repository

import (
	"github.com/huandu/go-sqlbuilder"
)

func (repository *CRUDRepository[Model]) DeleteReturnNull(where string) ([]Model, error) {
	builder := sqlbuilder.DeleteFrom(repository.table)
	builder.Where(where)

	query, args := builder.Build()

	result, err := repository.db.Exec(query, args)
	if err != nil {
		return nil, err
	}

	result.RowsAffected()

	return nil, nil
}
