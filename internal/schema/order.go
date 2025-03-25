package schema

import (
	"encoding/json"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
)

type Order[Model any] map[string]string

func (o *Order[Model]) UnmarshalText(text []byte) error {
	return json.Unmarshal(text, (*map[string]string)(o))
}

// Schema returns a schema representing this value on the wire.
// It returns the schema of the contained type.
func (o *Order[Model]) Schema(r huma.Registry) *huma.Schema {
	schema := &huma.Schema{
		Type:                 huma.TypeObject,
		Properties:           map[string]*huma.Schema{},
		AdditionalProperties: false,
	}

	modelType := reflect.TypeFor[Model]()
	for i := range modelType.NumField() {
		field := modelType.Field(i)
		schema.Properties[field.Tag.Get("json")] = &huma.Schema{
			Type: huma.TypeString,
			Enum: []any{"ASC", "DESC"},
		}
	}

	return schema
}
