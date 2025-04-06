package service

import (
	"context"
	"log/slog"

	"github.com/ckoliber/gocrud/internal/schema"
)

// DeleteBulkInput represents the input for the DeleteBulk operation
type DeleteBulkInput[Model any] struct {
	Where schema.Where[Model] `query:"where" doc:"Entity where" example:"{}"`
}

// DeleteBulkOutput represents the output for the DeleteBulk operation
type DeleteBulkOutput[Model any] struct {
	Body []Model
}

// DeleteBulk deletes multiple resources
func (s *CRUDService[Model]) DeleteBulk(ctx context.Context, i *DeleteBulkInput[Model]) (*DeleteBulkOutput[Model], error) {
	slog.Debug("Executing DeleteBulk operation", slog.Any("where", i.Where))

	// Execute BeforeDelete hook if defined
	if s.hooks.BeforeDelete != nil {
		if err := s.hooks.BeforeDelete(ctx, (*map[string]any)(&i.Where)); err != nil {
			slog.Error("BeforeDelete hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Delete the resources in the repository
	result, err := s.repo.Delete(ctx, (*map[string]any)(&i.Where))
	if err != nil {
		slog.Error("Failed to delete resources in DeleteBulk", slog.Any("error", err))
		return nil, err
	}

	// Execute AfterDelete hook if defined
	if s.hooks.AfterDelete != nil {
		if err := s.hooks.AfterDelete(ctx, &result); err != nil {
			slog.Error("AfterDelete hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed DeleteBulk operation", slog.Any("result", result))
	return &DeleteBulkOutput[Model]{
		Body: result,
	}, nil
}
