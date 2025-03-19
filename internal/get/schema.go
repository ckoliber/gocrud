package get

import "github.com/ckoliber/gocrud/internal/utils"

type InputOne[Model any] struct {
	ID string `path:"id" maxLength:"30" doc:"ID" example:"1"`
}

type OutputOne[Model any] struct {
	Body Model
}

type InputBulk[Model any] struct {
	Skip  int                `query:"skip" min:"0" doc:"Get skip" example:"0"`
	Limit int                `query:"limit" max:"100" doc:"Get limit" example:"50"`
	Order utils.Order[Model] `query:"order,deepObject" doc:"Get order" example:"{}"`
	Where utils.Where[Model] `query:"where,deepObject" doc:"Get where" example:"{}"`
}

type OutputBulk[Model any] struct {
	Body []Model
}
