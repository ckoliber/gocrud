package api

import "context"

type PutOneInput[Model any] struct {
	ID string `path:"id" maxLength:"30" doc:"ID"`
}
type PutOneOutput[Model any] struct {
	Body Model
}

func PutOne[Model any](ctx context.Context, i *PutOneInput[Model]) (*PutOneOutput[Model], error) {
	return nil, nil
}
