package schema

import (
	"encoding/json"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
)

type Order[Model any] map[string]string

func (o *Order[Model]) Validate() error {
	fields := map[string]string{}
	_type := reflect.TypeFor[Model]()
	for i := range _type.NumField() {
		field := _type.Field(i)
		fields[field.Tag.Get("json")] = ""
	}

	// for key, val := range *o {
	// 	if _, ok := fields[key]; !ok {
	// 		return huma.Error422UnprocessableEntity("invalid order key " + key)
	// 	}
	// 	if val != "ASC" && val != "DESC" {
	// 		return huma.Error422UnprocessableEntity("invalid order value " + key)
	// 	}
	// }

	return nil
}

func (o *Order[Model]) UnmarshalText(text []byte) error {
	if err := json.Unmarshal(text, (*map[string]string)(o)); err != nil {
		return err
	}

	if err := o.Validate(); err != nil {
		return err
	}

	return nil
}

func (o *Order[Model]) Schema(r huma.Registry) *huma.Schema {
	name := "Order" + huma.DefaultSchemaNamer(reflect.TypeFor[Model](), "")
	schema := &huma.Schema{
		Type: huma.TypeString,
	}

	r.Map()[name] = schema
	return schema
}
