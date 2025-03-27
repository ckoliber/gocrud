package service

import (
	"context"
)

type PostSingleInput[Model any] struct {
	Body Model
}
type PostSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) PostSingle(ctx context.Context, i *PostSingleInput[Model]) (*PostSingleOutput[Model], error) {
	if s.hooks.PreCreate != nil {
		if err := s.hooks.PreCreate(&[]Model{i.Body}); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Create(&[]Model{i.Body})
	if err != nil {
		return nil, err
	}

	if s.hooks.PostCreate != nil {
		if err := s.hooks.PostCreate(&result); err != nil {
			return nil, err
		}
	}

	return &PostSingleOutput[Model]{
		Body: result[0],
	}, nil
}
