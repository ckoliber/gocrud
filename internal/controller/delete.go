package controller

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type DeleteSingleInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type DeleteSingleOutput[Model any] struct {
	Body Model
}
type DeleteBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where" doc:"Get where" example:"{}"`
}
type DeleteBulkOutput[Model any] struct {
	Body []Model
}

func (controller *CRUDController[Model]) DeleteSingle(ctx context.Context, i *DeleteSingleInput[Model]) (*DeleteSingleOutput[Model], error) {
	return nil, nil
}

func (controller *CRUDController[Model]) DeleteBulk(ctx context.Context, i *DeleteBulkInput[Model]) (*DeleteBulkOutput[Model], error) {
	return nil, nil
}
