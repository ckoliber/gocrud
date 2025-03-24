package gocrud

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/huandu/go-sqlbuilder"

	"github.com/ckoliber/gocrud/internal/controller"
	"github.com/ckoliber/gocrud/internal/repository"
	"github.com/ckoliber/gocrud/internal/utils"
)

type Options[Model any] struct {
	Flavor sqlbuilder.Flavor

	ReadEnable   bool
	CreateEnable bool
	UpdateEnable bool
	DeleteEnable bool

	CreateBulk bool
	UpdateBulk bool
	DeleteBulk bool

	UpdateReturn int // Count - ID - Record
	DeleteReturn int // Count - ID - Record

	UpdatePartition bool
	DeletePartition bool

	MapRead   func(skip int, limit int, order map[string]string, where sqlbuilder.WhereClause, columns []string)
	MapCreate func(columns []string, models []Model)
	MapUpdate func(skip int, limit int, order map[string]string, where sqlbuilder.WhereClause, columns []string, models []Model)
	MapDelete func(skip int, limit int, order map[string]string, where sqlbuilder.WhereClause, columns []string)
}

func Register[Model any](api huma.API, db *sql.DB, options *Options[Model]) {
	_controller := controller.CRUDController[Model]{}
	_repository := repository.CRUDRepository[Model]{
		table:   utils.GetModelTable[Model](),
		columns: utils.GetModelColumns[Model](),
	}

	// Register GET operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("get-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET One %s", name),
		Description: fmt.Sprintf("GET One %s", name),
		// Tags:        []string{"GET", "One", name},
	}, crud.GetSingle)
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("get-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodGet,
		Summary:     fmt.Sprintf("GET Bulk %s", name),
		Description: fmt.Sprintf("GET Bulk %s", name),
		// Tags:        []string{"GET", "Bulk", name},
	}, crud.GetBulk)

	// Register PUT operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("put-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT One %s", name),
		Description: fmt.Sprintf("PUT One %s", name),
		// Tags:        []string{"PUT", "One", name},
	}, crud.PutSingle)
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("put-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPut,
		Summary:     fmt.Sprintf("PUT Bulk %s", name),
		Description: fmt.Sprintf("PUT Bulk %s", name),
		// Tags:        []string{"PUT", "Bulk", name},
	}, crud.PutBulk)

	// Register POST operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("post-one-%s", key),
		Path:        fmt.Sprintf("/%s/one", key),
		Method:      http.MethodPost,
		Summary:     fmt.Sprintf("POST One %s", name),
		Description: fmt.Sprintf("POST One %s", name),
		// Tags:        []string{"POST", "One", name},
	}, crud.PostSingle)
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("post-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPost,
		Summary:     fmt.Sprintf("POST Bulk %s", name),
		Description: fmt.Sprintf("POST Bulk %s", name),
		// Tags:        []string{"POST", "Bulk", name},
	}, crud.PostBulk)

	// Register PATCH operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("patch-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH One %s", name),
		Description: fmt.Sprintf("PATCH One %s", name),
		// Tags:        []string{"PATCH", "One", name},
	}, crud.PatchSingle)
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("patch-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodPatch,
		Summary:     fmt.Sprintf("PATCH Bulk %s", name),
		Description: fmt.Sprintf("PATCH Bulk %s", name),
		// Tags:        []string{"PATCH", "Bulk", name},
	}, crud.PatchBulk)

	// Register DELETE operations
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("delete-one-%s", key),
		Path:        fmt.Sprintf("/%s/{id}", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE One %s", name),
		Description: fmt.Sprintf("DELETE One %s", name),
		// Tags:        []string{"DELETE", "One", name},
	}, crud.DeleteSingle)
	huma.Register(api, huma.Operation{
		OperationID: fmt.Sprintf("delete-bulk-%s", key),
		Path:        fmt.Sprintf("/%s", key),
		Method:      http.MethodDelete,
		Summary:     fmt.Sprintf("DELETE Bulk %s", name),
		Description: fmt.Sprintf("DELETE Bulk %s", name),
		// Tags:        []string{"DELETE", "Bulk", name},
	}, crud.DeleteBulk)
}
