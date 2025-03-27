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
	if s.hooks.PreCreate != nil {
		if err := s.hooks.PreCreate(&i.Body); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Create(&i.Body)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostCreate != nil {
		if err := s.hooks.PostCreate(&result); err != nil {
			return nil, err
		}
	}

	return &PostBulkOutput[Model]{
		Body: result,
	}, nil
}
