package controller

import "context"

type PutOneInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}
type PutOneOutput[Model any] struct {
	Body Model
}

func PutOne[Model any](ctx context.Context, i *PutOneInput[Model]) (*PutOneOutput[Model], error) {
	return nil, nil
}

type PutBulkInput[Model any] struct {
	Body []Model
}
type PutBulkOutput[Model any] struct {
	Body []Model
}

func PutBulk[Model any](ctx context.Context, i *PutBulkInput[Model]) (*PutBulkOutput[Model], error) {
	return nil, nil
}
