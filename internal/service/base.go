package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ckoliber/gocrud/internal/repository"
	"github.com/ckoliber/gocrud/internal/schema"
)

type CRUDHooks[Model any] struct {
	PreRead   func(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int) error
	PreCreate func(fields *schema.Fields[Model], models *[]Model) error
	PreUpdate func(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int, model *Model) error
	PreDelete func(fields *schema.Fields[Model], where *schema.Where[Model], order *schema.Order[Model], limit *int, skip *int) error

	PostRead   func(models *[]Model) error
	PostCreate func(models *[]Model) error
	PostUpdate func(models *[]Model) error
	PostDelete func(models *[]Model) error
}

type CRUDService[Model any] struct {
	name  string
	path  string
	repo  *repository.CRUDRepository[Model]
	hooks *CRUDHooks[Model]
}

func NewCRUDService[Model any](repo *repository.CRUDRepository[Model], hooks *CRUDHooks[Model]) *CRUDService[Model] {
	_type := reflect.TypeFor[Model]()

	result := &CRUDService[Model]{
		name:  _type.Name(),
		path:  fmt.Sprintf("/%s", strings.ToLower(_type.Name())),
		repo:  repo,
		hooks: hooks,
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
