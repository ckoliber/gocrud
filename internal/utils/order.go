package utils

import "github.com/danielgtaylor/huma/v2"

type Order[Model any] map[string]string

// UnmarshalJSON unmarshals this value from JSON input.
func (order *Order[Model]) UnmarshalJSON(b []byte) error {
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
func (order Order[Model]) Schema(r huma.Registry) *huma.Schema {
	return &huma.Schema{
		OneOf: []*huma.Schema{
			{
				Type: huma.TypeObject,
				Properties: map[string]*huma.Schema{
					"foo": {
						Type: huma.TypeString,
						Extensions: map[string]any{
							"x-custom-thing": "abc123",
						},
					},
				},
			},
			{
				Type: huma.TypeObject,
				Properties: map[string]*huma.Schema{
					"bar": {
						Type: huma.TypeString,
						Extensions: map[string]any{
							"x-custom-thing": "abc123",
						},
					},
				},
			},
		},
	}
}

func (order *Order[Model]) ToSQL() string {
	return ""
}
