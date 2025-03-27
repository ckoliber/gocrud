package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PatchBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where,deepObject" doc:"Entity where" example:"{}"`
	Body  Model
}
type PatchBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) PatchBulk(ctx context.Context, i *PatchBulkInput[Model]) (o *PatchBulkOutput[Model], err error) {
	if s.hooks.PreUpdate != nil {
		if err := s.hooks.PreUpdate((*map[string]any)(&i.Where), &i.Body); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Update((*map[string]any)(&i.Where), &i.Body)
	if err != nil {
		return nil, err
	}

	if s.hooks.PostUpdate != nil {
		if err := s.hooks.PostUpdate(&o.Body); err != nil {
			return nil, err
		}
	}

	return &PatchBulkOutput[Model]{
		Body: result,
	}, nil
}
