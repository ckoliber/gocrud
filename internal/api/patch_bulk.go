package api

import (
	"context"
)

type PatchBulkInput[Model any] struct {
	Where Where[Model] `path:"where,deepObject" doc:"Where"`
}
type PatchBulkOutput[Model any] struct {
	Body []Model
}

func PatchBulk[Model any](ctx context.Context, i *PatchBulkInput[Model]) (*PatchBulkOutput[Model], error) {
	return nil, nil
}
