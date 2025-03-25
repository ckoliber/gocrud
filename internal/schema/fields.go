package schema

import (
	"encoding/json"

	"github.com/danielgtaylor/huma/v2"
)

type Fields[Model any] []string

func (f *Fields[Model]) UnmarshalText(text []byte) error {
	return json.Unmarshal(text, (*[]string)(f))
}

// Schema returns a schema representing this value on the wire.
// It returns the schema of the contained type.
func (f *Fields[Model]) Schema(r huma.Registry) *huma.Schema {
	schema := &huma.Schema{
		Type: huma.TypeArray,
		Items: &huma.Schema{
			Type: huma.TypeString,
			Enum: []any{"ASC", "DESC"},
		},
		AdditionalProperties: false,
	}

	return schema
}
