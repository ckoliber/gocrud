package controller

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type DeleteOneInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type DeleteOneOutput[Model any] struct {
	Body Model
}

func DeleteOne[Model any](ctx context.Context, i *DeleteOneInput[Model]) (*DeleteOneOutput[Model], error) {
	return nil, nil
}

type DeleteBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where" doc:"Get where" example:"{}"`
}
type DeleteBulkOutput[Model any] struct {
	Body []Model
}

func DeleteBulk[Model any](ctx context.Context, i *DeleteBulkInput[Model]) (*DeleteBulkOutput[Model], error) {
	return nil, nil
}
