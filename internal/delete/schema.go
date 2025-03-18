package delete

import "github.com/ckoliber/gocrud/internal/utils"

type InputOne[Model any] struct {
	ID string `path:"id" maxLength:"30" doc:"ID"`
}

type OutputOne[Model any] struct {
	Body Model
}

type InputBulk[Model any] struct {
	Where utils.Where[Model] `path:"where" doc:"Where"`
}

type OutputBulk[Model any] struct {
	Body []Model
}
