package repository

import "github.com/ckoliber/gocrud/internal/schema"

func (r *CRUDRepository[Model]) Delete(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int) ([]Model, error) {
	return nil, nil
}
