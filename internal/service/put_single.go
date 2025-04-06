package service

import (
	"context"
	"reflect"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
)

type PutSingleInput[Model any] struct {
	ID   string `path:"id" doc:"Entity identifier"`
	Body Model
}
type PutSingleOutput[Model any] struct {
	Body Model
}

func (s *CRUDService[Model]) PutSingle(ctx context.Context, i *PutSingleInput[Model]) (*PutSingleOutput[Model], error) {
	field := reflect.ValueOf(&i.Body).Elem().FieldByName(s.key)
	for field.Kind() == reflect.Pointer {
		field = field.Elem()
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := strconv.ParseInt(i.ID, 10, 64)
		if err != nil {
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		field.SetInt(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := strconv.ParseUint(i.ID, 10, 64)
		if err != nil {
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		field.SetUint(value)
	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(i.ID, 64)
		if err != nil {
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		field.SetFloat(value)
	case reflect.Complex64, reflect.Complex128:
		value, err := strconv.ParseComplex(i.ID, 128)
		if err != nil {
			return nil, huma.Error422UnprocessableEntity(err.Error())
		}
		field.SetComplex(value)
	case reflect.String:
		field.SetString(i.ID)
	default:
		return nil, huma.Error422UnprocessableEntity("invalid identifier type")
	}

	if s.hooks.BeforePut != nil {
		if err := s.hooks.BeforePut(ctx, &[]Model{i.Body}); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.Put(ctx, &[]Model{i.Body})
	if err != nil {
		return nil, err
	} else if len(result) <= 0 {
		return nil, huma.Error404NotFound("entity not found")
	}

	if s.hooks.AfterPut != nil {
		if err := s.hooks.AfterPut(ctx, &result); err != nil {
			return nil, err
		}
	}

	return &PutSingleOutput[Model]{
		Body: result[0],
	}, nil
}
