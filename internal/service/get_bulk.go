package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type GetBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where" doc:"Entity where" example:"{}"`
	Order schema.Order[Model] `query:"order" doc:"Entity order" example:"{}"`
	Limit int                 `query:"limit" min:"1" doc:"Entity limit" example:"50"`
	Skip  int                 `query:"skip" min:"0" doc:"Entity skip" example:"0"`
}
type GetBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) GetBulk(ctx context.Context, i *GetBulkInput[Model]) (*GetBulkOutput[Model], error) {
	if s.hooks.PreGet != nil {
		if err := s.hooks.PreGet(ctx, (*map[string]any)(&i.Where), (*map[string]any)(&i.Order), &i.Limit, &i.Skip); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Get(ctx, (*map[string]any)(&i.Where), (*map[string]any)(&i.Order), &i.Limit, &i.Skip)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostGet != nil {
		if err := s.hooks.PostGet(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &GetBulkOutput[Model]{
		Body: result,
	}, nil
}
