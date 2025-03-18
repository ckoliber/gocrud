package post

type InputOne[Model any] struct {
	Body Model
}

type OutputOne[Model any] struct {
	Body Model
}

type InputBulk[Model any] struct {
	Body []Model
}

type OutputBulk[Model any] struct {
	Body []Model
}
