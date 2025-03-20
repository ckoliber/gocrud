package api

import "context"

type PostOneInput[Model any] struct {
	ID string `path:"id" maxLength:"30" doc:"ID"`
}
type PostOneOutput[Model any] struct {
	Body Model
}

func PostOne[Model any](ctx context.Context, i *PostOneInput[Model]) (*PostOneOutput[Model], error) {
	return nil, nil
}
