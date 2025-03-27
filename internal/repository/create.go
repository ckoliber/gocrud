package repository

import (
	"github.com/ckoliber/gocrud/internal/schema"
)

func (r *CRUDRepository[Model]) Create(fields *schema.Fields[Model], models *[]Model) ([]Model, error) {
	return nil, nil
}
