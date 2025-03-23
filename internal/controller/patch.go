package controller

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PatchOneInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}
type PatchOneOutput[Model any] struct {
	Body Model
}

func PatchOne[Model any](ctx context.Context, i *PatchOneInput[Model]) (*PatchOneOutput[Model], error) {
	return nil, nil
}

type PatchBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where" doc:"Get where" example:"{}"`
	Body  Model
}
type PatchBulkOutput[Model any] struct {
	Body []Model
}

func PatchBulk[Model any](ctx context.Context, i *PatchBulkInput[Model]) (*PatchBulkOutput[Model], error) {
	return nil, nil
}
