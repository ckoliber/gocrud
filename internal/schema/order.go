package schema

import (
	"encoding/json"
	"errors"
	"log/slog"
	"reflect"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type Order[Model any] map[string]any

var orderRegistry huma.Registry

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

	_type := reflect.TypeFor[Model]()
	for i := range _type.NumField() {
		field := _type.Field(i)
		schema.Properties[strings.Split(field.Tag.Get("json"), ",")[0]] = &huma.Schema{
			Type: huma.TypeString,
			Enum: []any{"ASC", "DESC"},
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
