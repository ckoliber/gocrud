package repository

import (
	"github.com/huandu/go-sqlbuilder"
)

func (repository *CRUDRepository[Model]) Update(model Model, where string) ([]Model, error) {
	builder := sqlbuilder.Update(repository.table) //.Set(model)

	query, args := builder.Build()

	result, err := repository.db.Exec(query, args)
	if err != nil {
		return nil, err
	}

	result.RowsAffected()

	return nil, nil
}
