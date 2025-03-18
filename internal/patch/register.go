package patch

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API, path string) {
	huma.Register(api, huma.Operation{
		OperationID: "patch-one-...",
		Method:      http.MethodPatch,
		Path:        path + "/one",
		Summary:     "PATCH one ...",
		Description: "PATCH one ...",
		Tags:        []string{"PATCH", "One", "..."},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "patch-bulk-...",
		Method:      http.MethodPatch,
		Path:        path,
		Summary:     "PATCH bulk ...",
		Description: "PATCH bulk ...",
		Tags:        []string{"PATCH", "Bulk", "..."},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
