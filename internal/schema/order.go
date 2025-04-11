package schema

import (
	"encoding/json"
	"errors"
	"log/slog"
	"reflect"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

var orderRegistry huma.Registry

type Order[Model any] map[string]any

func (o *Order[Model]) UnmarshalText(text []byte) error {
	// Unmarshal the text into the Order map
	if err := json.Unmarshal(text, (*map[string]any)(o)); err != nil {
		slog.Error("Failed to unmarshal text into Order", slog.Any("error", err))
		return err
	}

	// Validate the unmarshaled data against the schema
	name := "Order" + huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := orderRegistry.Map()[name]
	result := huma.ValidateResult{}
	huma.Validate(orderRegistry, schema, huma.NewPathBuffer([]byte(""), 0), huma.ModeReadFromServer, (map[string]any)(*o), &result)
	if len(result.Errors) > 0 {
		slog.Error("Validation errors in Order", slog.Any("errors", result.Errors))
		return errors.Join(result.Errors...)
	}

	slog.Debug("Successfully unmarshaled and validated Order", slog.Any("order", *o))
	return nil
}

func (o *Order[Model]) Schema(r huma.Registry) *huma.Schema {
	// Generate and register the schema for the Order type
	name := "Order" + huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := &huma.Schema{
		Type:                 huma.TypeObject,
		Properties:           map[string]*huma.Schema{},
		AdditionalProperties: false,
	}

	// Add field-specific properties to the schema
	_type := reflect.TypeFor[Model]()
	for idx := range _type.NumField() {
		_field := _type.Field(idx)
		if _field.Name != "_" {
			if tag := _field.Tag.Get("json"); tag != "" {
				if _schema := o.FieldSchema(_field); _schema != nil {
					if tag != "-" {
						// primitive fields detected, name it with the json tag
						schema.Properties[strings.Split(tag, ",")[0]] = _schema
					}
				}
			}
		}
	}

	// Precompute messages and update the registry
	schema.PrecomputeMessages()
	r.Map()[name] = schema
	orderRegistry = r

	slog.Debug("Schema generated for Order", slog.String("name", name), slog.Any("schema", schema))
	return &huma.Schema{
		Type: huma.TypeString,
	}
}

func (o *Order[Model]) FieldSchema(field reflect.StructField) *huma.Schema {
	// Get the field deep inside array or slice or pointer types
	_field := field.Type
	for _field.Kind() == reflect.Array || _field.Kind() == reflect.Slice || _field.Kind() == reflect.Pointer {
		_field = _field.Elem()
	}

	switch _field.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		// For fields of primitive types, return a schema with enum values
		return &huma.Schema{
			Type: huma.TypeString,
			Enum: []any{"ASC", "DESC"},
		}
	}

	slog.Debug("Unsupported field type for Order", slog.Any("field", field))
	return nil
}
