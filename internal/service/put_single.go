package service

import (
	"context"
)

type PutSingleInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}
type PutSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) PutSingle(ctx context.Context, i *PutSingleInput[Model]) (*PutSingleOutput[Model], error) {
	if s.hooks.PrePut != nil {
		if err := s.hooks.PrePut(&[]Model{i.Body}); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Put(&[]Model{i.Body})
	if err != nil {
		return nil, err
	}

	if s.hooks.PostPut != nil {
		if err := s.hooks.PostPut(&result); err != nil {
			return nil, err
		}
	}

	return &PutSingleOutput[Model]{
		Body: result[0],
	}, nil
}
