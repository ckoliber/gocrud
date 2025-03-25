package gocrud

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"github.com/ckoliber/gocrud/internal/repository"
	"github.com/ckoliber/gocrud/internal/service"
)

type Mode int

const (
	BulkSingle Mode = iota
	Single
	None
)

type Options[Model any] struct {
	GETMode    Mode
	PUTMode    Mode
	POSTMode   Mode
	PATCHMode  Mode
	DELETEMode Mode

	service.CRUDHooks[Model]
}

func Register[Model any](api huma.API, db *sql.DB, options *Options[Model]) {
	repo := repository.NewCRUDRepository[Model](db)
	svc := service.NewCRUDService(repo, &options.CRUDHooks)

	// Register GET operations
	if options.GETMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("get-single-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath() + "/{id}",
			Method:      http.MethodGet,
			Summary:     fmt.Sprintf("GET Single %s", svc.GetName()),
			Description: fmt.Sprintf("GET Single %s", svc.GetName()),
			// Tags:        []string{"GET", "Single", svc.GetName()},
		}, svc.GetSingle)
	}
	if options.GETMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("get-bulk-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath(),
			Method:      http.MethodGet,
			Summary:     fmt.Sprintf("GET Bulk %s", svc.GetName()),
			Description: fmt.Sprintf("GET Bulk %s", svc.GetName()),
			// Tags:        []string{"GET", "Bulk", svc.GetName()},
		}, svc.GetBulk)
	}

	// Register PUT operations
	if options.PUTMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("put-single-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath() + "/{id}",
			Method:      http.MethodPut,
			Summary:     fmt.Sprintf("PUT Single %s", svc.GetName()),
			Description: fmt.Sprintf("PUT Single %s", svc.GetName()),
			// Tags:        []string{"PUT", "Single", svc.GetName()},
		}, svc.PutSingle)
	}
	if options.PUTMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("put-bulk-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath(),
			Method:      http.MethodPut,
			Summary:     fmt.Sprintf("PUT Bulk %s", svc.GetName()),
			Description: fmt.Sprintf("PUT Bulk %s", svc.GetName()),
			// Tags:        []string{"PUT", "Bulk", svc.GetName()},
		}, svc.PutBulk)
	}

	// Register POST operations
	if options.POSTMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("post-single-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath() + "/one",
			Method:      http.MethodPost,
			Summary:     fmt.Sprintf("POST Single %s", svc.GetName()),
			Description: fmt.Sprintf("POST Single %s", svc.GetName()),
			// Tags:        []string{"POST", "Single", name},
		}, svc.PostSingle)
	}
	if options.POSTMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("post-bulk-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath(),
			Method:      http.MethodPost,
			Summary:     fmt.Sprintf("POST Bulk %s", svc.GetName()),
			Description: fmt.Sprintf("POST Bulk %s", svc.GetName()),
			// Tags:        []string{"POST", "Bulk", svc.GetName()},
		}, svc.PostBulk)
	}

	// Register PATCH operations
	if options.PATCHMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("patch-single-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath() + "/{id}",
			Method:      http.MethodPatch,
			Summary:     fmt.Sprintf("PATCH Single %s", svc.GetName()),
			Description: fmt.Sprintf("PATCH Single %s", svc.GetName()),
			// Tags:        []string{"PATCH", "Single", svc.GetName()},
		}, svc.PatchSingle)
	}
	if options.PATCHMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("patch-bulk-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath(),
			Method:      http.MethodPatch,
			Summary:     fmt.Sprintf("PATCH Bulk %s", svc.GetName()),
			Description: fmt.Sprintf("PATCH Bulk %s", svc.GetName()),
			// Tags:        []string{"PATCH", "Bulk", svc.GetName()},
		}, svc.PatchBulk)
	}

	// Register DELETE operations
	if options.DELETEMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("delete-single-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath() + "/{id}",
			Method:      http.MethodDelete,
			Summary:     fmt.Sprintf("DELETE Single %s", svc.GetName()),
			Description: fmt.Sprintf("DELETE Single %s", svc.GetName()),
			// Tags:        []string{"DELETE", "Single", svc.GetName()},
		}, svc.DeleteSingle)
	}
	if options.DELETEMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("delete-bulk-%s", strings.ToLower(svc.GetName())),
			Path:        svc.GetPath(),
			Method:      http.MethodDelete,
			Summary:     fmt.Sprintf("DELETE Bulk %s", svc.GetName()),
			Description: fmt.Sprintf("DELETE Bulk %s", svc.GetName()),
			// Tags:        []string{"DELETE", "Bulk", svc.GetName()},
		}, svc.DeleteBulk)
	}
}
