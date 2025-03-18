package get

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/huandu/go-sqlbuilder"
)

func Register[Model any](api huma.API, path string) {
	huma.Register(api, huma.Operation{
		OperationID: "get-one-...",
		Path:        path + "/one",
		Method:      http.MethodGet,
		Summary:     "GET one ...",
		Description: "GET one ...",
		Tags:        []string{"GET", "One", "..."},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-bulk-...",
		Path:        path,
		Method:      http.MethodGet,
		Summary:     "GET bulk ...",
		Description: "GET bulk ...",
		Tags:        []string{"GET", "Bulk", "..."},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		sb := sqlbuilder.NewSelectBuilder()
		sb.Select("id", "name", "age").From("users")
		sb.OrderBy(i.Order.ToSQL())
		sb.Where(i.Where.ToSQL())
		sb.Limit(i.Limit)
		sb.Offset(i.Skip)

		sql, args := sb.Build()
		fmt.Println("--------------------")
		fmt.Println(sql)
		fmt.Println(args)
		fmt.Println("--------------------")

		return nil, errors.New(sql)
	})
}
