package utils

import (
	"reflect"
	"strings"
)

func GetModelPath[Model any]() *string {
	_type := reflect.TypeFor[Model]()

	if field, ok := _type.FieldByName("_"); ok {
		if value := field.Tag.Get("key"); value != "" {
			return &value
		}
	}

	return nil
}

func GetModelName[Model any]() string {
	_type := reflect.TypeFor[Model]()

	return _type.Name()
}

func GetModelTable[Model any]() string {
	model := reflect.TypeFor[Model]()
	name := model.Name()
	key := strings.ToLower(name)
	if metaField, ok := model.FieldByName("_"); ok {
		if value := metaField.Tag.Get("name"); value != "" {
			name = value
		}
		if value := metaField.Tag.Get("key"); value != "" {
			key = value
		}
	}
}

func GetModelColumns[Model any]() []string {
	_type := reflect.TypeFor[Model]()
}
