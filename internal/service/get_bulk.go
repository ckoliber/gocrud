package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type GetBulkInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	Where  schema.Where[Model]  `query:"where,deepObject" doc:"Entity where" example:"{}"`
	Order  schema.Order[Model]  `query:"order,deepObject" doc:"Entity order" example:"{}"`
	Limit  int                  `query:"limit" min:"1" doc:"Entity limit" example:"50"`
	Skip   int                  `query:"skip" min:"0" doc:"Entity skip" example:"0"`
}
type GetBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) GetBulk(ctx context.Context, i *GetBulkInput[Model]) (*GetBulkOutput[Model], error) {
	if s.hooks.PreRead != nil {
		if err := s.hooks.PreRead(&i.Fields, &i.Where, &i.Order, &i.Limit, &i.Skip); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Read(&i.Fields, &i.Where, &i.Order, &i.Limit, &i.Skip)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostRead != nil {
		if err := s.hooks.PostRead(&result); err != nil {
			return nil, err
		}
	}

	return &GetBulkOutput[Model]{
		Body: result,
	}, nil
}
