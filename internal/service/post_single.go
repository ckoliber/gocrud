package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PostSingleInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	Body   Model
}
type PostSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) PostSingle(ctx context.Context, i *PostSingleInput[Model]) (*PostSingleOutput[Model], error) {
	if s.hooks.PreCreate != nil {
		if err := s.hooks.PreCreate(&i.Fields, &[]Model{i.Body}); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Create(&i.Fields, &[]Model{i.Body})
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
