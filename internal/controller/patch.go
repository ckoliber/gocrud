package controller

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PatchSingleInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}
type PatchSingleOutput[Model any] struct {
	Body Model
}
type PatchBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where" doc:"Get where" example:"{}"`
	Body  Model
}
type PatchBulkOutput[Model any] struct {
	Body []Model
}

func (controller *CRUDController[Model]) PatchSingle(ctx context.Context, i *PatchSingleInput[Model]) (*PatchSingleOutput[Model], error) {
	return nil, nil
}

func (controller *CRUDController[Model]) PatchBulk(ctx context.Context, i *PatchBulkInput[Model]) (*PatchBulkOutput[Model], error) {
	return nil, nil
}
