package repository

import "github.com/ckoliber/gocrud/internal/schema"

func (r *CRUDRepository[Model]) Delete(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int) ([]Model, error) {
	builder := r.model.DeleteFrom(r.table)
	builder.Where(where)

	query, args := builder.Build()

	result, err := r.db.Exec(query, args)
	if err != nil {
		return nil, err
	}

	result.RowsAffected()

	return nil, nil
}
