package api

import "context"

type GetOneInput[Model any] struct {
	ID string `path:"id" maxLength:"30" doc:"ID"`
}
type GetOneOutput[Model any] struct {
	Body Model
}

func GetOne[Model any](ctx context.Context, i *GetOneInput[Model]) (*GetOneOutput[Model], error) {
	return nil, nil
}
