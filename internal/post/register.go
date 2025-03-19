package post

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
		OperationID: fmt.Sprintf("post-one-%s", key),
		Path:        fmt.Sprintf("/%s/one", key),
		Method:      http.MethodPost,
		Summary:     fmt.Sprintf("POST One %s", name),
		Description: fmt.Sprintf("POST One %s", name),
		Tags:        []string{"POST", "One", name},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("post-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPost,
		Summary:     fmt.Sprintf("POST Bulk %s", name),
		Description: fmt.Sprintf("POST Bulk %s", name),
		Tags:        []string{"POST", "Bulk", name},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
