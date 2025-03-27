package service

import (
	"context"
	"reflect"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PutBulkInput[Model any] struct {
	Body []Model
}
type PutBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) PutBulk(ctx context.Context, i *PutBulkInput[Model]) (*PutBulkOutput[Model], error) {
	o := &PutBulkOutput[Model]{}

	for _, model := range i.Body {
		where := schema.Where[Model]{s.id: reflect.Indirect(reflect.ValueOf(model)).FieldByName(s.key).String()}

		if s.hooks.PreUpdate != nil {
			if err := s.hooks.PreUpdate((*map[string]any)(&where), &model); err != nil {
				return nil, err
			}
		}

		result, err := s.repo.Update((*map[string]any)(&where), &model)
		if err != nil {
			return nil, err
		}

		if s.hooks.PostUpdate != nil {
			if err := s.hooks.PostUpdate(&result); err != nil {
				return nil, err
			}
		}

		o.Body = append(o.Body, result...)
	}

	return o, nil
}
