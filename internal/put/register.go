package put

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API, path string) {
	huma.Register(api, huma.Operation{
		OperationID: "put-one-...",
		Method:      http.MethodPut,
		Path:        path + "/one",
		Summary:     "PUT one ...",
		Description: "PUT one ...",
		Tags:        []string{"Create", "One", "..."},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "put-bulk-...",
		Method:      http.MethodPut,
		Path:        path,
		Summary:     "PUT bulk ...",
		Description: "PUT bulk ...",
		Tags:        []string{"PUT", "Bulk", "..."},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
