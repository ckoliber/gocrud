package service

import (
	"context"
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
)

type PostSingleInput[Model any] struct {
	Body Model
}
type PostSingleOutput[Model any] struct {
	Body Model
}

// PostSingle creates a single resource
func (s *CRUDService[Model]) PostSingle(ctx context.Context, i *PostSingleInput[Model]) (*PostSingleOutput[Model], error) {
	slog.Debug("Executing PostSingle operation", slog.Any("input", i))

	// Execute BeforePost hook if defined
	if s.hooks.BeforePost != nil {
		if err := s.hooks.BeforePost(ctx, &[]Model{i.Body}); err != nil {
			slog.Error("BeforePost hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Create the resource in the repository
	result, err := s.repo.Post(ctx, &[]Model{i.Body})
	if err != nil {
		slog.Error("Failed to create resource in PostSingle", slog.Any("error", err))
		return nil, err
	} else if len(result) <= 0 {
		slog.Error("Entity not found in PostSingle")
		return nil, huma.Error404NotFound("entity not found")
	}

	// Execute AfterPost hook if defined
	if s.hooks.AfterPost != nil {
		if err := s.hooks.AfterPost(ctx, &result); err != nil {
			slog.Error("AfterPost hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed PostSingle operation", slog.Any("result", result))
	return &PostSingleOutput[Model]{
		Body: result[0],
	}, nil
}
