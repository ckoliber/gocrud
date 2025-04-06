package service

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ckoliber/gocrud/internal/repository"
)

type CRUDHooks[Model any] struct {
	BeforeGet    func(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) error
	BeforePut    func(ctx context.Context, models *[]Model) error
	BeforePost   func(ctx context.Context, models *[]Model) error
	BeforeDelete func(ctx context.Context, where *map[string]any) error

	AfterGet    func(ctx context.Context, models *[]Model) error
	AfterPut    func(ctx context.Context, models *[]Model) error
	AfterPost   func(ctx context.Context, models *[]Model) error
	AfterDelete func(ctx context.Context, models *[]Model) error
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

	idField := _type.Field(0)
	if idField.Name == "_" {
		idField = _type.Field(1)
	}

	result := &CRUDService[Model]{
		id:    strings.Split(idField.Tag.Get("json"), ",")[0],
		key:   idField.Name,
		name:  strings.ToLower(_type.Name()),
		path:  fmt.Sprintf("/%s", strings.ToLower(_type.Name())),
		repo:  repo,
		hooks: hooks,
	}

	return result
}

func (s *CRUDService[Model]) GetName() string {
	return s.name
}

func (s *CRUDService[Model]) GetPath() string {
	return s.path
}
