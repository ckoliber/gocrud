package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type DeleteBulkInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	Where  schema.Where[Model]  `query:"where,deepObject" doc:"Entity where" example:"{}"`
	Order  schema.Order[Model]  `query:"order,deepObject" doc:"Entity order" example:"{}"`
	Limit  int                  `query:"limit" min:"1" doc:"Entity limit" example:"50"`
	Skip   int                  `query:"skip" min:"0" doc:"Entity skip" example:"0"`
}
type DeleteBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) DeleteBulk(ctx context.Context, i *DeleteBulkInput[Model]) (*DeleteBulkOutput[Model], error) {
	if err := s.hooks.PreDelete(&i.Fields, &i.Where, &i.Order, &i.Limit, &i.Skip); err != nil {
		return nil, err
	}

	result, err := s.repo.Delete(&i.Fields, &i.Where, &i.Order, &i.Limit, &i.Skip)
	if err != nil {
		return nil, err
	}

	if err := s.hooks.PostDelete(&result); err != nil {
		return nil, err
	}

	return &DeleteBulkOutput[Model]{
		Body: result,
	}, nil
}
