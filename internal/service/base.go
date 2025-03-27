package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ckoliber/gocrud/internal/repository"
)

type CRUDHooks[Model any] struct {
	PreRead   func(where *map[string]any, order *map[string]string, limit *int, skip *int) error
	PreUpdate func(where *map[string]any, model *Model) error
	PreDelete func(where *map[string]any) error
	PreCreate func(models *[]Model) error

	PostRead   func(models *[]Model) error
	PostUpdate func(models *[]Model) error
	PostDelete func(models *[]Model) error
	PostCreate func(models *[]Model) error
}

type CRUDService[Model any] struct {
	id    string
	key   string
	name  string
	path  string
	repo  repository.Repository[Model]
	hooks *CRUDHooks[Model]
}

func NewCRUDService[Model any](repo repository.Repository[Model], hooks *CRUDHooks[Model]) *CRUDService[Model] {
	_type := reflect.TypeFor[Model]()

	result := &CRUDService[Model]{
		id:    "id",
		key:   "ID",
		name:  _type.Name(),
		path:  fmt.Sprintf("/%s", strings.ToLower(_type.Name())),
		repo:  repo,
		hooks: hooks,
	}

	for i := range _type.NumField() {
		field := _type.Field(i)
		if val, ok := field.Tag.Lookup("id"); ok && val == "true" {
			result.id = field.Tag.Get("json")
			result.key = field.Name
		}
	}

	if field, ok := _type.FieldByName("_"); ok {
		if value := field.Tag.Get("name"); value != "" {
			result.name = value
		}
		if value := field.Tag.Get("path"); value != "" {
			result.path = value
		}
	}

	return result
}

func (s *CRUDService[Model]) GetName() string {
	return s.name
}

func (s *CRUDService[Model]) GetPath() string {
	return s.path
}
