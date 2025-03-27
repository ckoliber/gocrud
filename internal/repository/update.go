package repository

import "github.com/ckoliber/gocrud/internal/schema"

func (r *CRUDRepository[Model]) Update(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int, model *Model) ([]Model, error) {
	return nil, nil
}
