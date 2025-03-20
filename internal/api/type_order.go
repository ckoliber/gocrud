package api

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
)

type Order[Model any] map[string]string

func (order *Order[Model]) UnmarshalText(text []byte) error {
	return json.Unmarshal(text, (*map[string]string)(order))
}

// Schema returns a schema representing this value on the wire.
// It returns the schema of the contained type.
func (order *Order[Model]) Schema(r huma.Registry) *huma.Schema {
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

func (order *Order[Model]) ToSQL() (result []string) {
	for key, val := range *order {
		result = append(result, fmt.Sprintf("%s %s", key, val))
	}

	return result
}
