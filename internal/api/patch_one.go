package api

import "context"

type PatchOneInput[Model any] struct {
	ID string `path:"id" maxLength:"30" doc:"ID"`
}
type PatchOneOutput[Model any] struct {
	Body Model
}

func PatchOne[Model any](ctx context.Context, i *PatchOneInput[Model]) (*PatchOneOutput[Model], error) {
	return nil, nil
}
