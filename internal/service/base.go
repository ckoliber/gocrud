package service

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"github.com/ckoliber/gocrud/internal/repository"
)

// CRUDHooks defines hooks that can be executed before and after CRUD operations
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

// CRUDService provides CRUD operations for a given repository
type CRUDService[Model any] struct {
	id    string
	key   string
	name  string
	path  string
	repo  repository.Repository[Model]
	hooks *CRUDHooks[Model]
}

// NewCRUDService initializes a new CRUD service
func NewCRUDService[Model any](repo repository.Repository[Model], hooks *CRUDHooks[Model]) *CRUDService[Model] {
	// Reflect on the Model type to extract metadata
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

	slog.Debug("Initialized CRUDService", slog.String("name", result.name), slog.String("path", result.path), slog.String("id", result.id))
	return result
}

// GetName returns the name of the resource
func (s *CRUDService[Model]) GetName() string {
	slog.Debug("Fetching resource name", slog.String("name", s.name))
	return s.name
}

// GetPath returns the API path for the resource
func (s *CRUDService[Model]) GetPath() string {
	slog.Debug("Fetching resource path", slog.String("path", s.path))
	return s.path
}
