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
	for idx := range _type.NumField() {
		_field := _type.Field(idx)
		if _field.Name != "_" {
			if tag := _field.Tag.Get("json"); tag != "" {
				if _schema := w.FieldSchema(_field); _schema != nil {
					if tag == "-" {
						schema.Properties[_field.Tag.Get("db")] = _schema
					} else {
						schema.Properties[strings.Split(tag, ",")[0]] = _schema
					}
				}
			}
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

func (w *Where[Model]) FieldSchema(field reflect.StructField) *huma.Schema {
	_field := field.Type
	for _field.Kind() == reflect.Array || _field.Kind() == reflect.Slice || _field.Kind() == reflect.Pointer {
		_field = _field.Elem()
	}

	switch _field.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		result := &huma.Schema{
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

		if _method, ok := field.Type.MethodByName("Operations"); ok {
			var model Model
			value := reflect.ValueOf(model).FieldByName(field.Name)
			operations := _method.Func.Call([]reflect.Value{value})[0].Interface()
			for key := range operations.(map[string]func(string, ...string) string) {
				result.Properties[key] = &huma.Schema{
					Type:  huma.TypeString,
					Items: &huma.Schema{Type: huma.TypeString},
				}
			}
		}

		return result
	case reflect.Struct:
		name := "Where" + huma.DefaultSchemaNamer(_field, "")
		return &huma.Schema{
			Ref: "#/components/schemas/" + name,
		}
	}

	slog.Debug("Unsupported field type for Where", slog.Any("field", field))
	return nil
}
