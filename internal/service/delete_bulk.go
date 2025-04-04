package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type DeleteBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where" doc:"Entity where" example:"{}"`
}
type DeleteBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) DeleteBulk(ctx context.Context, i *DeleteBulkInput[Model]) (*DeleteBulkOutput[Model], error) {
	if s.hooks.PreDelete != nil {
		if err := s.hooks.PreDelete(ctx, (*map[string]any)(&i.Where)); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Delete(ctx, (*map[string]any)(&i.Where))
	if err != nil {
		return nil, err
	}

	if s.hooks.PostDelete != nil {
		if err := s.hooks.PostDelete(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &DeleteBulkOutput[Model]{
		Body: result,
	}, nil
}
