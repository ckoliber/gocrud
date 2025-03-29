package service

import (
	"context"
)

type PostBulkInput[Model any] struct {
	Body []Model
}
type PostBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) PostBulk(ctx context.Context, i *PostBulkInput[Model]) (*PostBulkOutput[Model], error) {
	if s.hooks.PrePost != nil {
		if err := s.hooks.PrePost(ctx, &i.Body); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Post(ctx, &i.Body)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostPost != nil {
		if err := s.hooks.PostPost(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &PostBulkOutput[Model]{
		Body: result,
	}, nil
}
