package get

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/huandu/go-sqlbuilder"
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
		OperationID: fmt.Sprintf("get-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET One %s", name),
		Description: fmt.Sprintf("GET One %s", name),
		Tags:        []string{"GET", "One", name},
	}, func(ctx context.Context, i *InputOne[Model]) (*OutputOne[Model], error) {
		return nil, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("get-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET Bulk %s", name),
		Description: fmt.Sprintf("GET Bulk %s", name),
		Tags:        []string{"GET", "Bulk", name},
	}, func(ctx context.Context, i *InputBulk[Model]) (*OutputBulk[Model], error) {
		sb := sqlbuilder.NewSelectBuilder()
		sb.Select("id", "name", "age").From("users")
		if sql, err := i.Order.ToSQL(); err == nil {
			sb.OrderBy(sql)
		}
		// if sql, err := i.Where.ToSQL(); err == nil {
		// 	sb.Where(sql)
		// }
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
