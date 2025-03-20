package api

import "context"

type DeleteBulkInput[Model any] struct {
	Where Where[Model] `query:"where" doc:"Get where" example:"{}"`
}
type DeleteBulkOutput[Model any] struct {
	Body []Model
}

func DeleteBulk[Model any](ctx context.Context, i *DeleteBulkInput[Model]) (*DeleteBulkOutput[Model], error) {
	return nil, nil
}
