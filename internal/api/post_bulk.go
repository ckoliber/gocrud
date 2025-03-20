package api

import (
	"context"
)

type PostBulkInput[Model any] struct {
	Where Where[Model] `path:"where,deepObject" doc:"Where"`
}
type PostBulkOutput[Model any] struct {
	Body []Model
}

func PostBulk[Model any](ctx context.Context, i *PostBulkInput[Model]) (*PostBulkOutput[Model], error) {
	return nil, nil
}
