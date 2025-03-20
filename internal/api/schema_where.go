package api

import (
	"fmt"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
)

type Where[Model any] map[string]any

func (where *Where[Model]) Resolve(ctx huma.Context) error {
	fmt.Println("Resolve")
	return nil
}

// Schema returns a schema representing this value on the wire.
// It returns the schema of the contained type.
func (where *Where[Model]) Schema(r huma.Registry) *huma.Schema {
	name := "Where" + huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := &huma.Schema{
		// Ref:        "#/components/schemas/" + name,
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

	modelType := reflect.TypeFor[Model]()
	for i := range modelType.NumField() {
		field := modelType.Field(i)
		schema.Properties[field.Tag.Get("json")] = &huma.Schema{
			Type: huma.TypeString,
			Enum: []any{"ASC", "DESC"},
		}
	}

	r.Map()[name] = schema
	return schema
}

func (where *Where[Model]) ToSQL() (string, error) {
	fmt.Println(where)
	return "", nil
}
