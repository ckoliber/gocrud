package utils

import (
	"fmt"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
)

type Where[Model any] map[string]any

// Schema returns a schema representing this value on the wire.
// It returns the schema of the contained type.
func (where *Where[Model]) Schema(r huma.Registry) *huma.Schema {
	ref := "#/components/schemas/Where"
	schema := &huma.Schema{
		Ref: ref,
		OneOf: []*huma.Schema{
			{
				Type:                 huma.TypeObject,
				Properties:           map[string]*huma.Schema{},
				AdditionalProperties: false,
			},
			{
				Type: huma.TypeObject,
				Properties: map[string]*huma.Schema{
					"_not": {
						Ref: ref,
					},
				},
			},
			{
				Type: huma.TypeObject,
				Properties: map[string]*huma.Schema{
					"_and": {
						Type: huma.TypeArray,
						Items: &huma.Schema{
							Ref: ref,
						},
					},
				},
			},
			{
				Type: huma.TypeObject,
				Properties: map[string]*huma.Schema{
					"_or": {
						Type: huma.TypeArray,
						Items: &huma.Schema{
							Ref: ref,
						},
					},
				},
			},
		},
	}

	modelType := reflect.TypeFor[Model]()
	for i := range modelType.NumField() {
		field := modelType.Field(i)
		schema.OneOf[0].Properties[field.Tag.Get("json")] = &huma.Schema{
			OneOf: []*huma.Schema{
				{
					Type:                 huma.TypeObject,
					Properties:           map[string]*huma.Schema{"_lt": {Type: huma.TypeString}},
					AdditionalProperties: false,
				},
				{
					Type:                 huma.TypeObject,
					Properties:           map[string]*huma.Schema{"_gt": {Type: huma.TypeString}},
					AdditionalProperties: false,
				},
				{
					Type:                 huma.TypeObject,
					Properties:           map[string]*huma.Schema{"_lte": {Type: huma.TypeString}},
					AdditionalProperties: false,
				},
				{
					Type:                 huma.TypeObject,
					Properties:           map[string]*huma.Schema{"_gte": {Type: huma.TypeString}},
					AdditionalProperties: false,
				},
			},
		}
	}

	return schema
}

func (where *Where[Model]) ToSQL() (string, error) {
	fmt.Println(where)
	return "", nil
}
