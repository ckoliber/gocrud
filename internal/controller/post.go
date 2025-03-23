package controller

import "context"

type PostOneInput[Model any] struct {
	Body Model
}
type PostOneOutput[Model any] struct {
	Body Model
}

func PostOne[Model any](ctx context.Context, i *PostOneInput[Model]) (*PostOneOutput[Model], error) {
	return nil, nil
}

type PostBulkInput[Model any] struct {
	Body []Model
}
type PostBulkOutput[Model any] struct {
	Body []Model
}

func PostBulk[Model any](ctx context.Context, i *PostBulkInput[Model]) (*PostBulkOutput[Model], error) {
	return nil, nil
}
