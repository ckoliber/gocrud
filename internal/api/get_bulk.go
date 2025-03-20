package api

import (
	"context"
)

type GetBulkInput[Model any] struct {
	Where Where[Model] `path:"where,deepObject" doc:"Where"`
}
type GetBulkOutput[Model any] struct {
	Body []Model
}

func GetBulk[Model any](ctx context.Context, i *GetBulkInput[Model]) (*GetBulkOutput[Model], error) {
	return nil, nil
}
