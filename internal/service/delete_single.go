package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type DeleteSingleInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type DeleteSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) DeleteSingle(ctx context.Context, i *DeleteSingleInput[Model]) (*DeleteSingleOutput[Model], error) {
	where := schema.Where[Model]{s.id: map[string]any{"_eq": i.ID}}

	if s.hooks.BeforeDelete != nil {
		if err := s.hooks.BeforeDelete(ctx, (*map[string]any)(&where)); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Delete(ctx, (*map[string]any)(&where))
	if err != nil {
		return nil, err
	}

	if s.hooks.AfterDelete != nil {
		if err := s.hooks.AfterDelete(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &DeleteSingleOutput[Model]{
		Body: result[0],
	}, nil
}
