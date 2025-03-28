package schema

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
)

type Where[Model any] map[string]any

var whereRegistry = huma.NewMapRegistry("#/where/", huma.DefaultSchemaNamer)

func (w *Where[Model]) UnmarshalText(text []byte) error {
	if err := json.Unmarshal(text, (*map[string]any)(w)); err != nil {
		return err
	}

	name := huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := whereRegistry.Map()[name]
	result := huma.ValidateResult{}
	huma.Validate(whereRegistry, schema, huma.NewPathBuffer([]byte(""), 0), huma.ModeReadFromServer, (map[string]any)(*w), &result)
	if len(result.Errors) > 0 {
		return errors.Join(result.Errors...)
	}

	return nil
}

func (w *Where[Model]) Schema(r huma.Registry) *huma.Schema {
	name := huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := &huma.Schema{
		Type: huma.TypeObject,
		Properties: map[string]*huma.Schema{
			"_not": {
				Ref: "#/where/" + name,
			},
			"_and": {
				Type: huma.TypeArray,
				Items: &huma.Schema{
					Ref: "#/where/" + name,
				},
			},
			"_or": {
				Type: huma.TypeArray,
				Items: &huma.Schema{
					Ref: "#/where/" + name,
				},
			},
		},
		AdditionalProperties: false,
	}

	_type := reflect.TypeFor[Model]()
	for i := range _type.NumField() {
		field := _type.Field(i)
		schema.Properties[field.Tag.Get("json")] = &huma.Schema{
			Type: huma.TypeObject,
			Properties: map[string]*huma.Schema{
				"_eq":  {Type: huma.TypeString},
				"_neq": {Type: huma.TypeString},
				"_gt":  {Type: huma.TypeString},
				"_gte": {Type: huma.TypeString},
				"_lt":  {Type: huma.TypeString},
				"_lte": {Type: huma.TypeString},
			},
			AdditionalProperties: false,
		}
	}

	schema.PrecomputeMessages()
	whereRegistry.Map()[name] = schema

	return &huma.Schema{
		Type: huma.TypeString,
	}
}
