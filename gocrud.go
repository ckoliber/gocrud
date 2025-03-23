package gocrud

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"github.com/ckoliber/gocrud/internal/controller"
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

	// Register GET operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("get-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET One %s", name),
		Description: fmt.Sprintf("GET One %s", name),
		// Tags:        []string{"GET", "One", name},
	}, controller.GetOne[Model])
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("get-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET Bulk %s", name),
		Description: fmt.Sprintf("GET Bulk %s", name),
		// Tags:        []string{"GET", "Bulk", name},
	}, controller.GetBulk[Model])

	// Register PUT operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("put-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT One %s", name),
		Description: fmt.Sprintf("PUT One %s", name),
		// Tags:        []string{"PUT", "One", name},
	}, controller.PutOne[Model])
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("put-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT Bulk %s", name),
		Description: fmt.Sprintf("PUT Bulk %s", name),
		// Tags:        []string{"PUT", "Bulk", name},
	}, controller.PutBulk[Model])

	// Register POST operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("post-one-%s", key),
		Path:        fmt.Sprintf("/%s/one", key),
		Method:      http.MethodPost,
		Summary:     fmt.Sprintf("POST One %s", name),
		Description: fmt.Sprintf("POST One %s", name),
		// Tags:        []string{"POST", "One", name},
	}, controller.PostOne[Model])
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("post-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPost,
		Summary:     fmt.Sprintf("POST Bulk %s", name),
		Description: fmt.Sprintf("POST Bulk %s", name),
		// Tags:        []string{"POST", "Bulk", name},
	}, controller.PostBulk[Model])

	// Register PATCH operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("patch-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH One %s", name),
		Description: fmt.Sprintf("PATCH One %s", name),
		// Tags:        []string{"PATCH", "One", name},
	}, controller.PatchOne[Model])
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("patch-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH Bulk %s", name),
		Description: fmt.Sprintf("PATCH Bulk %s", name),
		// Tags:        []string{"PATCH", "Bulk", name},
	}, controller.PatchBulk[Model])

	// Register DELETE operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("delete-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE One %s", name),
		Description: fmt.Sprintf("DELETE One %s", name),
		// Tags:        []string{"DELETE", "One", name},
	}, controller.DeleteOne[Model])
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("delete-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE Bulk %s", name),
		Description: fmt.Sprintf("DELETE Bulk %s", name),
		// Tags:        []string{"DELETE", "Bulk", name},
	}, controller.DeleteBulk[Model])
}
