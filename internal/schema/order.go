package schema

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type Order[Model any] map[string]any

var orderRegistry huma.Registry

func (o *Order[Model]) UnmarshalText(text []byte) error {
	if err := json.Unmarshal(text, (*map[string]any)(o)); err != nil {
		return err
	}

	name := "Order" + huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := orderRegistry.Map()[name]
	result := huma.ValidateResult{}
	huma.Validate(orderRegistry, schema, huma.NewPathBuffer([]byte(""), 0), huma.ModeReadFromServer, (map[string]any)(*o), &result)
	if len(result.Errors) > 0 {
		return errors.Join(result.Errors...)
	}

	return nil
}

func (o *Order[Model]) Schema(r huma.Registry) *huma.Schema {
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

	schema.PrecomputeMessages()
	r.Map()[name] = schema
	orderRegistry = r

	return &huma.Schema{
		Type: huma.TypeString,
	}
}
