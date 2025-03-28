package service

import (
	"context"
)

type PutBulkInput[Model any] struct {
	Body []Model
}
type PutBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) PutBulk(ctx context.Context, i *PutBulkInput[Model]) (*PutBulkOutput[Model], error) {
	if s.hooks.PrePut != nil {
		if err := s.hooks.PrePut(&i.Body); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Put(&i.Body)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostPut != nil {
		if err := s.hooks.PostPut(&result); err != nil {
			return nil, err
		}
	}

	return &PutBulkOutput[Model]{
		Body: result,
	}, nil
}
