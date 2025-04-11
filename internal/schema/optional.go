package schema

import (
	"reflect"

	"github.com/danielgtaylor/huma/v2"
)

type Optional[Type any] struct {
	Value Type
	IsSet bool
}

// Define schema to use wrapped type
func (o *Optional[Type]) Schema(r huma.Registry) *huma.Schema {
	return huma.SchemaFromType(r, reflect.TypeOf(o.Value))
}

// Expose wrapped value to receive parsed value from Huma
func (o *Optional[Type]) Receiver() reflect.Value {
	return reflect.ValueOf(o).Elem().Field(0)
}

// React to request param being parsed to update internal state
func (o *Optional[Type]) OnParamSet(isSet bool, parsed any) {
	o.IsSet = isSet
}

// Get the value of the wrapped type
func (o *Optional[Type]) Addr() *Type {
	if o.IsSet {
		return &o.Value
	}

	return nil
}
