package service

import (
	"context"
	"log/slog"
)

type PutBulkInput[Model any] struct {
	Body []Model
}
type PutBulkOutput[Model any] struct {
	Body []Model
}

// PutBulk updates multiple resources
func (s *CRUDService[Model]) PutBulk(ctx context.Context, i *PutBulkInput[Model]) (*PutBulkOutput[Model], error) {
	slog.Debug("Executing PutBulk operation", slog.Any("input", i))

	// Execute BeforePut hook if defined
	if s.hooks.BeforePut != nil {
		if err := s.hooks.BeforePut(ctx, &i.Body); err != nil {
			slog.Error("BeforePut hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Update the resources in the repository
	result, err := s.repo.Put(ctx, &i.Body)
	if err != nil {
		slog.Error("Failed to update resources in PutBulk", slog.Any("error", err))
		return nil, err
	}

	// Execute AfterPut hook if defined
	if s.hooks.AfterPut != nil {
		if err := s.hooks.AfterPut(ctx, &result); err != nil {
			slog.Error("AfterPut hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed PutBulk operation", slog.Any("result", result))
	return &PutBulkOutput[Model]{
		Body: result,
	}, nil
}
