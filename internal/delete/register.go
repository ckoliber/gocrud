package delete

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API, path string) {
	huma.Register(api, huma.Operation{
		OperationID: "delete-one-...",
		Method:      http.MethodDelete,
		Path:        path + "/one",
		Summary:     "DELETE one ...",
		Description: "DELETE one ...",
		Tags:        []string{"DELETE", "One", "..."},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-bulk-...",
		Method:      http.MethodDelete,
		Path:        path,
		Summary:     "DELETE bulk ...",
		Description: "DELETE bulk ...",
		Tags:        []string{"DELETE", "Bulk", "..."},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
