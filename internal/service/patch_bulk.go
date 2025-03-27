package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PatchBulkInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	Where  schema.Where[Model]  `query:"where,deepObject" doc:"Entity where" example:"{}"`
	Order  schema.Order[Model]  `query:"order,deepObject" doc:"Entity order" example:"{}"`
	Limit  int                  `query:"limit" min:"1" doc:"Entity limit" example:"50"`
	Skip   int                  `query:"skip" min:"0" doc:"Entity skip" example:"0"`
	Body   Model
}
type PatchBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) PatchBulk(ctx context.Context, i *PatchBulkInput[Model]) (o *PatchBulkOutput[Model], err error) {
	if s.hooks.PreUpdate != nil {
		if err := s.hooks.PreUpdate(&i.Fields, &i.Where, &i.Order, &i.Limit, &i.Skip, &i.Body); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Update(&i.Fields, &i.Where, &i.Order, &i.Limit, &i.Skip, &i.Body)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostUpdate != nil {
		if err := s.hooks.PostUpdate(&o.Body); err != nil {
			return nil, err
		}
	}

	return &PatchBulkOutput[Model]{
		Body: result,
	}, nil
}
