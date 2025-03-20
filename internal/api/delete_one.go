package api

import "context"

type DeleteOneInput[Model any] struct {
	ID string `path:"id" maxLength:"30" doc:"ID"`
}
type DeleteOneOutput[Model any] struct {
	Body Model
}

func DeleteOne[Model any](ctx context.Context, i *DeleteOneInput[Model]) (*DeleteOneOutput[Model], error) {
	return nil, nil
}
