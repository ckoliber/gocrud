package api

import "context"

type GetBulkInput[Model any] struct {
	Skip  int          `query:"skip" min:"0" doc:"Get skip" example:"0"`
	Limit int          `query:"limit" max:"100" doc:"Get limit" example:"50"`
	Order Order[Model] `query:"order" doc:"Get order" example:"{}"`
	Where Where[Model] `query:"where" doc:"Get where" example:"{}"`
}
type GetBulkOutput[Model any] struct {
	Body []Model
}

func GetBulk[Model any](ctx context.Context, i *GetBulkInput[Model]) (*GetBulkOutput[Model], error) {
	return nil, nil
}
