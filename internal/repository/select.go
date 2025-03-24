package repository

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

func (repository *CRUDRepository[Model]) Select(model Model, where string) ([]Model, error) {
	builder := sqlbuilder.Select(repository.columns...).From(repository.table)

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
