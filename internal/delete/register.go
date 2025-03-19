package delete

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
		OperationID: fmt.Sprintf("delete-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE One %s", name),
		Description: fmt.Sprintf("DELETE One %s", name),
		Tags:        []string{"DELETE", "One", name},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("delete-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE Bulk %s", name),
		Description: fmt.Sprintf("DELETE Bulk %s", name),
		Tags:        []string{"DELETE", "Bulk", name},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
