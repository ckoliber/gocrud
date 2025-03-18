package utils

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

type Where[Model any] struct {
	Value map[string]string
}

// UnmarshalJSON unmarshals this value from JSON input.
func (where *Where[Model]) UnmarshalJSON(b []byte) error {
	// if len(b) > 0 {
	// 	o.Sent = true
	// 	if bytes.Equal(b, []byte("null")) {
	// 		o.Null = true
	// 		return nil
	// 	}
	// 	return json.Unmarshal(b, &o.Value)
	// }
	return nil
}

// Schema returns a schema representing this value on the wire.
// It returns the schema of the contained type.
func (where Where[Model]) Schema(r huma.Registry) *huma.Schema {
	// schema := r.Schema(reflect.TypeOf(where.Value), true, "")
	// Build schema for comparison operators
	// comparisonSchema := &huma.Schema{
	// 	Type: huma.TypeObject,
	// 	Properties: map[string]*huma.Schema{
	// 		"eq":  {Type: huma.TypeString},
	// 		"neq": {Type: huma.TypeString},
	// 		"gt":  {Type: huma.TypeNumber},
	// 		"gte": {Type: huma.TypeNumber},
	// 		"lt":  {Type: huma.TypeNumber},
	// 		"lte": {Type: huma.TypeNumber},
	// 		"in": {
	// 			Type: huma.TypeArray,
	// 			Items: &huma.Schema{
	// 				Type: huma.TypeString,
	// 			},
	// 		},
	// 	},
	// }

	schema := &huma.Schema{
		OneOf: []*huma.Schema{
			{
				Type: huma.TypeObject,
				Properties: map[string]*huma.Schema{
					"foo": {Type: huma.TypeString},
				},
			},
			{
				Type: huma.TypeArray,
				Items: &huma.Schema{
					Type: huma.TypeObject,
					Properties: map[string]*huma.Schema{
						"foo": {Type: huma.TypeString},
					},
				},
			},
		},
	}

	fmt.Println(schema.Ref)

	// schema.OneOf = append(schema.OneOf,
	// 	&huma.Schema{

	// 	&huma.Schema{
	// 		Type: huma.TypeObject,
	// 		Properties: map[string]*huma.Schema{
	// 			"_or": {
	// 				Type:  huma.TypeArray,
	// 				Items: schema,
	// 			},
	// 		},
	// 	},
	// 	&huma.Schema{
	// 		Type: huma.TypeObject,
	// 		Properties: map[string]*huma.Schema{
	// 			"_and": {
	// 				Type:  huma.TypeArray,
	// 				Items: schema,
	// 			},
	// 		},
	// 	},
	// 	&huma.Schema{
	// 		Type: huma.TypeObject,
	// 		Properties: map[string]*huma.Schema{
	// 			"_not": schema,
	// 		},
	// 	},
	// )

	return schema
}

func (where *Where[Model]) ToSQL() string {
	return ""
}
