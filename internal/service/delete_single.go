package service

import (
	"context"
	"log/slog"

	"github.com/ckoliber/gocrud/internal/schema"
	"github.com/danielgtaylor/huma/v2"
)

type DeleteSingleInput[Model any] struct {
	ID string `path:"id" doc:"Entity identifier"`
}
type DeleteSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) DeleteSingle(ctx context.Context, i *DeleteSingleInput[Model]) (*DeleteSingleOutput[Model], error) {
	slog.Debug("Executing DeleteSingle operation", slog.String("id", i.ID))

	where := schema.Where[Model]{s.id: map[string]any{"_eq": i.ID}}

	// Execute BeforeDelete hook if defined
	if s.hooks.BeforeDelete != nil {
		if err := s.hooks.BeforeDelete(ctx, (*map[string]any)(&where)); err != nil {
			slog.Error("BeforeDelete hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Delete the resource in the repository
	result, err := s.repo.Delete(ctx, (*map[string]any)(&where))
	if err != nil {
		slog.Error("Failed to delete resource in DeleteSingle", slog.Any("error", err))
		return nil, err
	} else if len(result) <= 0 {
		slog.Warn("Entity not found in DeleteSingle", slog.String("id", i.ID))
		return nil, huma.Error404NotFound("entity not found")
	}

	// Execute AfterDelete hook if defined
	if s.hooks.AfterDelete != nil {
		if err := s.hooks.AfterDelete(ctx, &result); err != nil {
			slog.Error("AfterDelete hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed DeleteSingle operation", slog.Any("result", result[0]))
	return &DeleteSingleOutput[Model]{
		Body: result[0],
	}, nil
}
