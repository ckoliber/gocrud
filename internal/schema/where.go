package schema

import (
	"encoding/json"
	"errors"
	"log/slog"
	"reflect"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type Where[Model any] map[string]any

var whereRegistry huma.Registry

func (w *Where[Model]) UnmarshalText(text []byte) error {
	// Unmarshal the text into the Where map
	if err := json.Unmarshal(text, (*map[string]any)(w)); err != nil {
		slog.Error("Failed to unmarshal text into Where", slog.Any("error", err))
		return err
	}

	// Validate the unmarshaled data against the schema
	name := "Where" + huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := whereRegistry.Map()[name]
	result := huma.ValidateResult{}
	huma.Validate(whereRegistry, schema, huma.NewPathBuffer([]byte(""), 0), huma.ModeReadFromServer, (map[string]any)(*w), &result)
	if len(result.Errors) > 0 {
		slog.Error("Validation errors in Where", slog.Any("errors", result.Errors))
		return errors.Join(result.Errors...)
	}

	slog.Debug("Successfully unmarshaled and validated Where", slog.Any("where", *w))
	return nil
}

func (w *Where[Model]) Schema(r huma.Registry) *huma.Schema {
	// Generate and register the schema for the Where type
	name := "Where" + huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := &huma.Schema{
		Type: huma.TypeObject,
		Properties: map[string]*huma.Schema{
			"_not": {
				Ref: "#/components/schemas/" + name,
			},
			"_and": {
				Type: huma.TypeArray,
				Items: &huma.Schema{
					Ref: "#/components/schemas/" + name,
				},
			},
			"_or": {
				Type: huma.TypeArray,
				Items: &huma.Schema{
					Ref: "#/components/schemas/" + name,
				},
			},
		},
		AdditionalProperties: false,
	}

	// Add field-specific properties to the schema
	_type := reflect.TypeFor[Model]()
	for i := range _type.NumField() {
		field := _type.Field(i)
		schema.Properties[strings.Split(field.Tag.Get("json"), ",")[0]] = &huma.Schema{
			Type: huma.TypeObject,
			Properties: map[string]*huma.Schema{
				"_eq":     {Type: huma.TypeString},
				"_neq":    {Type: huma.TypeString},
				"_gt":     {Type: huma.TypeString},
				"_gte":    {Type: huma.TypeString},
				"_lt":     {Type: huma.TypeString},
				"_lte":    {Type: huma.TypeString},
				"_like":   {Type: huma.TypeString},
				"_nlike":  {Type: huma.TypeString},
				"_ilike":  {Type: huma.TypeString},
				"_nilike": {Type: huma.TypeString},
				"_in":     {Type: huma.TypeArray, Items: &huma.Schema{Type: huma.TypeString}},
				"_nin":    {Type: huma.TypeArray, Items: &huma.Schema{Type: huma.TypeString}},
			},
			AdditionalProperties: false,
		}
	}

	// Precompute messages and update the registry
	schema.PrecomputeMessages()
	r.Map()[name] = schema
	whereRegistry = r

	slog.Debug("Schema generated for Where", slog.String("name", name), slog.Any("schema", schema))
	return &huma.Schema{
		Type: huma.TypeString,
	}
}
