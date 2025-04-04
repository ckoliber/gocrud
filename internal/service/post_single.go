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
	if s.hooks.PrePost != nil {
		if err := s.hooks.PrePost(ctx, &[]Model{i.Body}); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Post(ctx, &[]Model{i.Body})
	if err != nil {
		return nil, err
	}

	if s.hooks.PostPost != nil {
		if err := s.hooks.PostPost(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &PostSingleOutput[Model]{
		Body: result[0],
	}, nil
}
