package service

import (
	"context"
	"log/slog"

	"github.com/ckoliber/gocrud/internal/schema"
)

// GetBulkInput defines the input parameters for the GetBulk operation
type GetBulkInput[Model any] struct {
	Where schema.Where[Model]  `query:"where" doc:"Entity where" example:"{}"`
	Order schema.Order[Model]  `query:"order" doc:"Entity order" example:"{}"`
	Limit schema.Optional[int] `query:"limit" min:"1" doc:"Entity limit" example:"50"`
	Skip  schema.Optional[int] `query:"skip" min:"0" doc:"Entity skip" example:"0"`
}

// GetBulkOutput defines the output structure for the GetBulk operation
type GetBulkOutput[Model any] struct {
	Body []Model
}

// GetBulk retrieves multiple resources with filtering and pagination
func (s *CRUDService[Model]) GetBulk(ctx context.Context, i *GetBulkInput[Model]) (*GetBulkOutput[Model], error) {
	slog.Debug("Executing GetBulk operation", slog.Any("where", i.Where), slog.Any("order", i.Order), slog.Any("limit", i.Limit), slog.Any("skip", i.Skip))

	// Execute BeforeGet hook if defined
	if s.hooks.BeforeGet != nil {
		if err := s.hooks.BeforeGet(ctx, i.Where.Addr(), i.Order.Addr(), i.Limit.Addr(), i.Skip.Addr()); err != nil {
			slog.Error("BeforeGet hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Fetch resources from the repository
	result, err := s.repo.Get(ctx, i.Where.Addr(), i.Order.Addr(), i.Limit.Addr(), i.Skip.Addr())
	if err != nil {
		slog.Error("Failed to fetch resources in GetBulk", slog.Any("error", err))
		return nil, err
	}

	// Execute AfterGet hook if defined
	if s.hooks.AfterGet != nil {
		if err := s.hooks.AfterGet(ctx, &result); err != nil {
			slog.Error("AfterGet hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed GetBulk operation", slog.Any("result", result))
	return &GetBulkOutput[Model]{
		Body: result,
	}, nil
}
