package gocrud

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/ckoliber/gocrud/internal/api"
	"github.com/danielgtaylor/huma/v2"
)

func Register[Model any](_api huma.API) {
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
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("get-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET One %s", name),
		Description: fmt.Sprintf("GET One %s", name),
		Tags:        []string{"GET", "One", name},
	}, api.GetOne[Model])
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("get-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET Bulk %s", name),
		Description: fmt.Sprintf("GET Bulk %s", name),
		Tags:        []string{"GET", "Bulk", name},
	}, api.GetBulk[Model])

	// Register PUT operations
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("put-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT One %s", name),
		Description: fmt.Sprintf("PUT One %s", name),
		Tags:        []string{"PUT", "One", name},
	}, api.PutOne[Model])
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("put-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT Bulk %s", name),
		Description: fmt.Sprintf("PUT Bulk %s", name),
		Tags:        []string{"PUT", "Bulk", name},
	}, api.PutBulk[Model])

	// Register POST operations
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("post-one-%s", key),
		Path:        fmt.Sprintf("/%s/one", key),
		Method:      http.MethodPost,
		Summary:     fmt.Sprintf("POST One %s", name),
		Description: fmt.Sprintf("POST One %s", name),
		Tags:        []string{"POST", "One", name},
	}, api.PostOne[Model])
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("post-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("POST Bulk %s", name),
		Description: fmt.Sprintf("POST Bulk %s", name),
		Tags:        []string{"POST", "Bulk", name},
	}, api.PostBulk[Model])

	// Register PATCH operations
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("patch-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH One %s", name),
		Description: fmt.Sprintf("PATCH One %s", name),
		Tags:        []string{"PATCH", "One", name},
	}, api.PatchOne[Model])
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("patch-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH Bulk %s", name),
		Description: fmt.Sprintf("PATCH Bulk %s", name),
		Tags:        []string{"PATCH", "Bulk", name},
	}, api.PatchBulk[Model])

	// Register DELETE operations
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("delete-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE One %s", name),
		Description: fmt.Sprintf("DELETE One %s", name),
		Tags:        []string{"DELETE", "One", name},
	}, api.DeleteOne[Model])
	huma.Register(_api, huma.Operation{
		OperationID: fmt.Sprintf("delete-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE Bulk %s", name),
		Description: fmt.Sprintf("DELETE Bulk %s", name),
		Tags:        []string{"DELETE", "Bulk", name},
	}, api.DeleteBulk[Model])
}
