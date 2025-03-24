package controller

import "context"

type PutSingleInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}
type PutSingleOutput[Model any] struct {
	Body Model
}
type PutBulkInput[Model any] struct {
	Body []Model
}
type PutBulkOutput[Model any] struct {
	Body []Model
}

func (controller *CRUDController[Model]) PutSingle(ctx context.Context, i *PutSingleInput[Model]) (*PutSingleOutput[Model], error) {
	return nil, nil
}

func (controller *CRUDController[Model]) PutBulk(ctx context.Context, i *PutBulkInput[Model]) (*PutBulkOutput[Model], error) {
	return nil, nil
}
