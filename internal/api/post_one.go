package api

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
