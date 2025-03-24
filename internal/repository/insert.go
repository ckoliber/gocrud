package repository

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

func (repository *CRUDRepository[Model]) Insert(models []Model) ([]Model, error) {
	builder := sqlbuilder.InsertInto(repository.table).Cols(repository.columns...)
	builder.Values(models).Returning(repository.columns...)

	query, args := builder.Build()

	rows, err := repository.db.Query(query, args)
	if err != nil {
		return nil, err
	}

	var result []Model
	rows.Scan(result)

	fmt.Println(query)
	fmt.Println(args)

	return result, nil
}
