package api

import "context"

type PatchBulkInput[Model any] struct {
	Where Where[Model] `query:"where" doc:"Get where" example:"{}"`
	Body  Model
}
type PatchBulkOutput[Model any] struct {
	Body []Model
}

func PatchBulk[Model any](ctx context.Context, i *PatchBulkInput[Model]) (*PatchBulkOutput[Model], error) {
	return nil, nil
}
