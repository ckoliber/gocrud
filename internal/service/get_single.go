package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
	"github.com/danielgtaylor/huma/v2"
)

type GetSingleInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type GetSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) GetSingle(ctx context.Context, i *GetSingleInput[Model]) (*GetSingleOutput[Model], error) {
	where := schema.Where[Model]{s.id: map[string]any{"_eq": i.ID}}

	if s.hooks.BeforeGet != nil {
		if err := s.hooks.BeforeGet(ctx, (*map[string]any)(&where), nil, nil, nil); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Get(ctx, (*map[string]any)(&where), nil, nil, nil)
	if err != nil {
		return nil, err
	} else if len(result) <= 0 {
		return nil, huma.Error404NotFound("entity not found")
	}

	if s.hooks.AfterGet != nil {
		if err := s.hooks.AfterGet(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &GetSingleOutput[Model]{
		Body: result[0],
	}, nil
}
