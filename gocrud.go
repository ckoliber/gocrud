package gocrud

import (
	"github.com/ckoliber/gocrud/internal/get"
	"github.com/ckoliber/gocrud/internal/put"
	"github.com/ckoliber/gocrud/internal/post"
	"github.com/ckoliber/gocrud/internal/patch"
	"github.com/ckoliber/gocrud/internal/delete"

	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API, path string) {
	get.Register[Model](api, path)
	put.Register[Model](api, path)
	post.Register[Model](api, path)
	patch.Register[Model](api, path)
	delete.Register[Model](api, path)
}
