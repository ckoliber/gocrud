package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ckoliber/gocrud/internal/schema"
)

type GetSingleInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type GetSingleOutput[Model any] struct {
	Body Model
}
type GetBulkInput[Model any] struct {
	Skip  int                 `query:"skip" min:"0" doc:"Get skip" example:"0"`
	Limit int                 `query:"limit" min:"1" doc:"Get limit" example:"50"`
	Order schema.Order[Model] `query:"order,deepObject" doc:"Get order" example:"{}"`
	Where schema.Where[Model] `query:"where,deepObject" doc:"Get where" example:"{}"`
}
type GetBulkOutput[Model any] struct {
	Body []Model
}

func (controller *CRUDController[Model]) GetSingle(ctx context.Context, i *GetSingleInput[Model]) (*GetSingleOutput[Model], error) {
	return nil, nil
}

func (controller *CRUDController[Model]) GetBulk(ctx context.Context, i *GetBulkInput[Model]) (*GetBulkOutput[Model], error) {
	data, _ := json.Marshal(i.Where)
	fmt.Println(string(data))
	return nil, nil
}
