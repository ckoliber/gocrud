package gocrud

import (
	"database/sql"
	"fmt"
	"net/http"

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
	GetMode    Mode
	PutMode    Mode
	PostMode   Mode
	DeleteMode Mode

	service.CRUDHooks[Model]
}

func Register[Model any](api huma.API, repo repository.Repository[Model], options *Options[Model]) {
	svc := service.NewCRUDService(repo, &options.CRUDHooks)

	// Register Get operations
	if options.GetMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("get-single-%s", svc.GetName()),
			Path:        svc.GetPath() + "/{id}",
			Method:      http.MethodGet,
			Summary:     fmt.Sprintf("Get single-%s", svc.GetName()),
			Description: fmt.Sprintf("Retrieves a single %s by its unique identifier. Returns full resource representation.", svc.GetName()),
		}, svc.GetSingle)
	}
	if options.GetMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("get-bulk-%s", svc.GetName()),
			Path:        svc.GetPath(),
			Method:      http.MethodGet,
			Summary:     fmt.Sprintf("Get bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Returns a paginated list of %s resources. Supports filtering, sorting and pagination parameters.", svc.GetName()),
		}, svc.GetBulk)
	}

	// Register Put operations
	if options.PutMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("put-single-%s", svc.GetName()),
			Path:        svc.GetPath() + "/one",
			Method:      http.MethodPut,
			Summary:     fmt.Sprintf("Put single-%s", svc.GetName()),
			Description: fmt.Sprintf("Full update operation for a %s resource. Requires complete resource representation.", svc.GetName()),
		}, svc.PutSingle)
	}
	if options.PutMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("put-bulk-%s", svc.GetName()),
			Path:        svc.GetPath(),
			Method:      http.MethodPut,
			Summary:     fmt.Sprintf("Put bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Batch update operation for multiple %s resources. Each resource requires complete representation.", svc.GetName()),
		}, svc.PutBulk)
	}

	// Register Post operations
	if options.PostMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("post-single-%s", svc.GetName()),
			Path:        svc.GetPath() + "/one",
			Method:      http.MethodPost,
			Summary:     fmt.Sprintf("Post single-%s", svc.GetName()),
			Description: fmt.Sprintf("Creates a new %s resource. Returns the created resource with generated identifier.", svc.GetName()),
		}, svc.PostSingle)
	}
	if options.PostMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("post-bulk-%s", svc.GetName()),
			Path:        svc.GetPath(),
			Method:      http.MethodPost,
			Summary:     fmt.Sprintf("Post bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Batch creation operation for multiple %s resources. Returns created resources with generated identifiers.", svc.GetName()),
		}, svc.PostBulk)
	}

	// Register Delete operations
	if options.DeleteMode <= Single {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("delete-single-%s", svc.GetName()),
			Path:        svc.GetPath() + "/{id}",
			Method:      http.MethodDelete,
			Summary:     fmt.Sprintf("Delete single-%s", svc.GetName()),
			Description: fmt.Sprintf("Permanently removes a %s resource by its identifier. This operation cannot be undone.", svc.GetName()),
		}, svc.DeleteSingle)
	}
	if options.DeleteMode <= BulkSingle {
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("delete-bulk-%s", svc.GetName()),
			Path:        svc.GetPath(),
			Method:      http.MethodDelete,
			Summary:     fmt.Sprintf("Delete bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Batch deletion operation for multiple %s resources. This operation cannot be undone.", svc.GetName()),
		}, svc.DeleteBulk)
	}
}

func NewRepository[Model any](db *sql.DB) repository.Repository[Model] {
	return repository.NewSQLRepository[Model](db)
}
