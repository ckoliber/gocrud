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
	if s.hooks.BeforePost != nil {
		if err := s.hooks.BeforePost(ctx, &i.Body); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Post(ctx, &i.Body)
	if err != nil {
		return nil, err
	}

	if s.hooks.AfterPost != nil {
		if err := s.hooks.AfterPost(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &PostBulkOutput[Model]{
		Body: result,
	}, nil
}
