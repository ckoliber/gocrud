package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type GetSingleInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type GetSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) GetSingle(ctx context.Context, i *GetSingleInput[Model]) (*GetSingleOutput[Model], error) {
	where := schema.Where[Model]{s.id: map[string]any{"_eq": i.ID}}

	if s.hooks.PreRead != nil {
		if err := s.hooks.PreRead((*map[string]any)(&where), nil, nil, nil); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Read((*map[string]any)(&where), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostRead != nil {
		if err := s.hooks.PostRead(&result); err != nil {
			return nil, err
		}
	}

	return &GetSingleOutput[Model]{
		Body: result[0],
	}, nil
}
