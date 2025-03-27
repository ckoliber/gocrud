package repository

import (
	"os"

	"github.com/ckoliber/gocrud/internal/schema"
)

func (r *CRUDRepository[Model]) Read(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int) ([]Model, error) {
	data := map[string]any{
		"action": "read",
		"fields": fields,
		"where":  where,
		"order":  order,
		"limit":  limit,
		"skip":   skip,
	}

	if err := r.template.Execute(os.Stdout, data); err != nil {
		return nil, err
	}

	return nil, nil
}
