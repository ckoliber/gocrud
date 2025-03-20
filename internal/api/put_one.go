package api

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
