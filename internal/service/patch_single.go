package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PatchSingleInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	ID     string               `path:"id" doc:"Entity identifier"`
	Body   Model
}
type PatchSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) PatchSingle(ctx context.Context, i *PatchSingleInput[Model]) (*PatchSingleOutput[Model], error) {
	where := schema.Where[Model]{"id": i.ID}

	if err := s.hooks.PreUpdate(&i.Fields, &where, nil, nil, nil, &i.Body); err != nil {
		return nil, err
	}

	result, err := s.repo.Update(&i.Fields, &where, nil, nil, nil, &i.Body)
	if err != nil {
		return nil, err
	}

	if err := s.hooks.PostUpdate(&result); err != nil {
		return nil, err
	}

	return &PatchSingleOutput[Model]{
		Body: result[0],
	}, nil
}
