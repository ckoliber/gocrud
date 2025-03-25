package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PostBulkInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	Body   []Model
}
type PostBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) PostBulk(ctx context.Context, i *PostBulkInput[Model]) (*PostBulkOutput[Model], error) {
	if err := s.hooks.PreCreate(&i.Fields, &i.Body); err != nil {
		return nil, err
	}

	result, err := s.repo.Create(&i.Fields, &i.Body)
	if err != nil {
		return nil, err
	}

	if err := s.hooks.PostCreate(&result); err != nil {
		return nil, err
	}

	return &PostBulkOutput[Model]{
		Body: result,
	}, nil
}
