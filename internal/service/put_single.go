package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PutSingleInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}
type PutSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) PutSingle(ctx context.Context, i *PutSingleInput[Model]) (*PutSingleOutput[Model], error) {
	where := schema.Where[Model]{s.id: i.ID}

	if s.hooks.PreUpdate != nil {
		if err := s.hooks.PreUpdate((*map[string]any)(&where), &i.Body); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Update((*map[string]any)(&where), &i.Body)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostUpdate != nil {
		if err := s.hooks.PostUpdate(&result); err != nil {
			return nil, err
		}
	}

	return &PutSingleOutput[Model]{
		Body: result[0],
	}, nil
}
