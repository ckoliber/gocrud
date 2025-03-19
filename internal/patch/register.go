package patch

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
		OperationID: fmt.Sprintf("patch-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH One %s", name),
		Description: fmt.Sprintf("PATCH One %s", name),
		Tags:        []string{"PATCH", "One", name},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("patch-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH Bulk %s", name),
		Description: fmt.Sprintf("PATCH Bulk %s", name),
		Tags:        []string{"PATCH", "Bulk", name},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
