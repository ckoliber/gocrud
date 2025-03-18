package post

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](api huma.API, path string) {
	huma.Register(api, huma.Operation{
		OperationID: "post-one-...",
		Method:      http.MethodPost,
		Path:        path + "/one",
		Summary:     "POST one ...",
		Description: "POST one ...",
		Tags:        []string{"POST", "One", "..."},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "post-bulk-...",
		Method:      http.MethodPost,
		Path:        path,
		Summary:     "POST bulk ...",
		Description: "POST bulk ...",
		Tags:        []string{"POST", "Bulk", "..."},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		return nil, nil
	})
}
