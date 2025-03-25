package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type DeleteSingleInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	ID     string               `path:"id" doc:"Entity identifier"`
}
type DeleteSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) DeleteSingle(ctx context.Context, i *DeleteSingleInput[Model]) (*DeleteSingleOutput[Model], error) {
	where := schema.Where[Model]{"id": i.ID}

	if err := s.hooks.PreDelete(&i.Fields, &where, nil, nil, nil); err != nil {
		return nil, err
	}

	result, err := s.repo.Delete(&i.Fields, &where, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := s.hooks.PostDelete(&result); err != nil {
		return nil, err
	}

	return &DeleteSingleOutput[Model]{
		Body: result[0],
	}, nil
}
