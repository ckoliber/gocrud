package patch

import "github.com/ckoliber/gocrud/utils"

type InputOne[Model any] struct {
	ID   string `path:"id" maxLength:"30" doc:"ID"`
	Body Model
}

type OutputOne[Model any] struct {
	Body Model
}

type InputBulk[Model any] struct {
	Where utils.Where[Model] `path:"where" doc:"Where"`
	Body  Model
}

type OutputBulk[Model any] struct {
	Body []Model
}
