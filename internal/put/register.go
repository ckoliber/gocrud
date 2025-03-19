package put

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API) {
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

	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("put-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT One %s", name),
		Description: fmt.Sprintf("PUT One %s", name),
		Tags:        []string{"PUT", "One", name},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("put-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT Bulk %s", name),
		Description: fmt.Sprintf("PUT Bulk %s", name),
		Tags:        []string{"PUT", "Bulk", name},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
