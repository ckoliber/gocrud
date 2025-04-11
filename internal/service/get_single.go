package service

import (
	"context"
	"log/slog"

	"github.com/ckoliber/gocrud/internal/schema"
	"github.com/danielgtaylor/huma/v2"
)

type GetSingleInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type GetSingleOutput[Model any] struct {
	Body Model
}

// GetSingle retrieves a single resource by its ID
func (s *CRUDService[Model]) GetSingle(ctx context.Context, i *GetSingleInput[Model]) (*GetSingleOutput[Model], error) {
	slog.Debug("Executing GetSingle operation", slog.String("id", i.ID))

	// Define the where clause for the get operation
	where := schema.Where[Model]{s.id: map[string]any{"_eq": i.ID}}

	// Execute BeforeGet hook if defined
	if s.hooks.BeforeGet != nil {
		if err := s.hooks.BeforeGet(ctx, (*map[string]any)(&where), nil, nil, nil); err != nil {
			slog.Error("BeforeGet hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Fetch the resource from the repository
	result, err := s.repo.Get(ctx, (*map[string]any)(&where), nil, nil, nil)
	if err != nil {
		slog.Error("Failed to fetch resource in GetSingle", slog.Any("error", err))
		return nil, err
	} else if len(result) <= 0 {
		slog.Error("Entity not found in GetSingle", slog.String("id", i.ID))
		return nil, huma.Error404NotFound("entity not found")
	}

	// Execute AfterGet hook if defined
	if s.hooks.AfterGet != nil {
		if err := s.hooks.AfterGet(ctx, &result); err != nil {
			slog.Error("AfterGet hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed GetSingle operation", slog.Any("result", result))
	return &GetSingleOutput[Model]{
		Body: result[0],
	}, nil
}
