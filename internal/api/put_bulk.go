package api

import (
	"context"
)

type PutBulkInput[Model any] struct {
	Where Where[Model] `path:"where,deepObject" doc:"Where"`
}
type PutBulkOutput[Model any] struct {
	Body []Model
}

func PutBulk[Model any](ctx context.Context, i *PutBulkInput[Model]) (*PutBulkOutput[Model], error) {
	return nil, nil
}
