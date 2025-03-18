package gocrud

import (
	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API, path string) {
	fmt.Println(path)
}
