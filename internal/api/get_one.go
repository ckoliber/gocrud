package api

import "context"

type GetOneInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type GetOneOutput[Model any] struct {
	Body Model
}

func GetOne[Model any](ctx context.Context, i *GetOneInput[Model]) (*GetOneOutput[Model], error) {
	return nil, nil
}
