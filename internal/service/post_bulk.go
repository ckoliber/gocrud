package service

import (
	"context"
	"log/slog"
)

type PostBulkInput[Model any] struct {
	Body []Model
}
type PostBulkOutput[Model any] struct {
	Body []Model
}

// PostBulk creates multiple resources
func (s *CRUDService[Model]) PostBulk(ctx context.Context, i *PostBulkInput[Model]) (*PostBulkOutput[Model], error) {
	slog.Debug("Executing PostBulk operation", slog.Any("input", i))

	// Execute BeforePost hook if defined
	if s.hooks.BeforePost != nil {
		if err := s.hooks.BeforePost(ctx, &i.Body); err != nil {
			slog.Error("BeforePost hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Create resources in the repository
	result, err := s.repo.Post(ctx, &i.Body)
	if err != nil {
		slog.Error("Failed to create resources in PostBulk", slog.Any("error", err))
		return nil, err
	}

	// Execute AfterPost hook if defined
	if s.hooks.AfterPost != nil {
		if err := s.hooks.AfterPost(ctx, &result); err != nil {
			slog.Error("AfterPost hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed PostBulk operation", slog.Any("result", result))
	return &PostBulkOutput[Model]{
		Body: result,
	}, nil
}
