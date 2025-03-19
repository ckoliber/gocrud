package gocrud

import (
	"github.com/ckoliber/gocrud/internal/delete"
	"github.com/ckoliber/gocrud/internal/get"
	"github.com/ckoliber/gocrud/internal/patch"
	"github.com/ckoliber/gocrud/internal/post"
	"github.com/ckoliber/gocrud/internal/put"

	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API) {
	get.Register[Model](api)
	put.Register[Model](api)
	post.Register[Model](api)
	patch.Register[Model](api)
	delete.Register[Model](api)
}
