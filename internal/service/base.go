package service

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ckoliber/gocrud/internal/repository"
)

type CRUDHooks[Model any] struct {
	PreGet    func(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) error
	PrePut    func(ctx context.Context, models *[]Model) error
	PrePost   func(ctx context.Context, models *[]Model) error
	PreDelete func(ctx context.Context, where *map[string]any) error

	PostGet    func(ctx context.Context, models *[]Model) error
	PostPut    func(ctx context.Context, models *[]Model) error
	PostPost   func(ctx context.Context, models *[]Model) error
	PostDelete func(ctx context.Context, models *[]Model) error
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
		if value := field.Tag.Get("id"); value == "true" {
			result.id = field.Tag.Get("json")
			result.key = field.Name
		}
	}

	if field, ok := _type.FieldByName("_"); ok {
		if value := field.Tag.Get("json"); value != "" {
			result.name = value
		}
		if value := field.Tag.Get("path"); value != "" {
			result.path = value
		}
	}

	return result
}

func (s *CRUDService[Model]) GetName() string {
	return strings.ToLower(s.name)
}

func (s *CRUDService[Model]) GetPath() string {
	return s.path
}
