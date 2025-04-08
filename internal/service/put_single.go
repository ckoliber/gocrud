package service

import (
	"context"
	"log/slog"
	"reflect"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
)

// PutSingleInput represents the input for the PutSingle operation
type PutSingleInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}

// PutSingleOutput represents the output for the PutSingle operation
type PutSingleOutput[Model any] struct {
	Body Model
}

// PutSingle updates a single resource
func (s *CRUDService[Model]) PutSingle(ctx context.Context, i *PutSingleInput[Model]) (*PutSingleOutput[Model], error) {
	slog.Debug("Executing PutSingle operation", slog.String("id", i.ID), slog.Any("body", i.Body))

	_field := reflect.ValueOf(&i.Body).Elem().FieldByName(s.key)
	for _field.Kind() == reflect.Pointer {
		_field = _field.Elem()
	}

	switch _field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := strconv.ParseInt(i.ID, 10, 64)
		if err != nil {
			slog.Error("Failed to parse ID as integer", slog.Any("error", err))
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		_field.SetInt(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := strconv.ParseUint(i.ID, 10, 64)
		if err != nil {
			slog.Error("Failed to parse ID as unsigned integer", slog.Any("error", err))
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		_field.SetUint(value)
	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(i.ID, 64)
		if err != nil {
			slog.Error("Failed to parse ID as float", slog.Any("error", err))
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		_field.SetFloat(value)
	case reflect.Complex64, reflect.Complex128:
		value, err := strconv.ParseComplex(i.ID, 128)
		if err != nil {
			slog.Error("Failed to parse ID as complex number", slog.Any("error", err))
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		_field.SetComplex(value)
	case reflect.String:
		_field.SetString(i.ID)
	default:
		slog.Error("Invalid identifier type", slog.String("id", i.ID))
		return nil, huma.Error422UnprocessableEntity("invalid identifier type")
	}

	// Execute BeforePut hook if defined
	if s.hooks.BeforePut != nil {
		if err := s.hooks.BeforePut(ctx, &[]Model{i.Body}); err != nil {
			slog.Error("BeforePut hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	// Update the resource in the repository
	result, err := s.repo.Put(ctx, &[]Model{i.Body})
	if err != nil {
		slog.Error("Failed to update resource in PutSingle", slog.Any("error", err))
		return nil, err
	} else if len(result) <= 0 {
		slog.Error("Entity not found", slog.String("id", i.ID))
		return nil, huma.Error404NotFound("entity not found")
	}

	// Execute AfterPut hook if defined
	if s.hooks.AfterPut != nil {
		if err := s.hooks.AfterPut(ctx, &result); err != nil {
			slog.Error("AfterPut hook failed", slog.Any("error", err))
			return nil, err
		}
	}

	slog.Debug("Successfully executed PutSingle operation", slog.Any("result", result[0]))
	return &PutSingleOutput[Model]{
		Body: result[0],
	}, nil
}
