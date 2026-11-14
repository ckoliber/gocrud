package gocrud

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"

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

type Config[Model any] struct {
	GetMode    Mode
	PutMode    Mode
	PostMode   Mode
	DeleteMode Mode

	BeforeGet    func(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) error
	BeforePut    func(ctx context.Context, models *[]Model) error
	BeforePost   func(ctx context.Context, models *[]Model) error
	BeforeDelete func(ctx context.Context, where *map[string]any) error

	AfterGet    func(ctx context.Context, models *[]Model) error
	AfterPut    func(ctx context.Context, models *[]Model) error
	AfterPost   func(ctx context.Context, models *[]Model) error
	AfterDelete func(ctx context.Context, models *[]Model) error
}

// Register sets up CRUD operations for the given API and repository based on the provided configuration.
func Register[Model any](api huma.API, repo repository.Repository[Model], config *Config[Model]) {
	// Initialize CRUD service with hooks
	svc := service.NewCRUDService(repo, &service.CRUDHooks[Model]{
		BeforeGet:    config.BeforeGet,
		BeforePut:    config.BeforePut,
		BeforePost:   config.BeforePost,
		BeforeDelete: config.BeforeDelete,
		AfterGet:     config.AfterGet,
		AfterPut:     config.AfterPut,
		AfterPost:    config.AfterPost,
		AfterDelete:  config.AfterDelete,
	})

	// Get path for operations
	path := svc.GetPath()

	// Register Get operations
	if config.GetMode <= Single {
		slog.Debug("Registering GetSingle operation", slog.String("path", path+"/{id}"))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("get-single-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Get single-%s", svc.GetName()),
			Description: fmt.Sprintf("Retrieves a single %s by its unique identifier. Returns full resource representation.", svc.GetName()),
			Path:        path + "/{id}",
			Method:      http.MethodGet,
		}, svc.GetSingle)
	}
	if config.GetMode <= BulkSingle {
		slog.Debug("Registering GetBulk operation", slog.String("path", path))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("get-bulk-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Get bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Returns a paginated list of %s resources. Supports filtering, sorting and pagination parameters.", svc.GetName()),
			Path:        path,
			Method:      http.MethodGet,
		}, svc.GetBulk)
	}

	// Register Put operations
	if config.PutMode <= Single {
		slog.Debug("Registering PutSingle operation", slog.String("path", path+"/{id}"))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("put-single-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Put single-%s", svc.GetName()),
			Description: fmt.Sprintf("Full update operation for a %s resource. Requires complete resource representation.", svc.GetName()),
			Path:        path + "/{id}",
			Method:      http.MethodPut,
		}, svc.PutSingle)
	}
	if config.PutMode <= BulkSingle {
		slog.Debug("Registering PutBulk operation", slog.String("path", path))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("put-bulk-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Put bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Batch update operation for multiple %s resources. Each resource requires complete representation.", svc.GetName()),
			Path:        path,
			Method:      http.MethodPut,
		}, svc.PutBulk)
	}

	// Register Post operations
	if config.PostMode <= Single {
		slog.Debug("Registering PostSingle operation", slog.String("path", path+"/one"))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("post-single-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Post single-%s", svc.GetName()),
			Description: fmt.Sprintf("Creates a new %s resource. Returns the created resource with generated identifier.", svc.GetName()),
			Path:        path + "/one",
			Method:      http.MethodPost,
		}, svc.PostSingle)
	}
	if config.PostMode <= BulkSingle {
		slog.Debug("Registering PostBulk operation", slog.String("path", path))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("post-bulk-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Post bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Batch creation operation for multiple %s resources. Returns created resources with generated identifiers.", svc.GetName()),
			Path:        path,
			Method:      http.MethodPost,
		}, svc.PostBulk)
	}

	// Register Delete operations
	if config.DeleteMode <= Single {
		slog.Debug("Registering DeleteSingle operation", slog.String("path", path+"/{id}"))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("delete-single-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Delete single-%s", svc.GetName()),
			Description: fmt.Sprintf("Permanently removes a %s resource by its identifier. This operation cannot be undone.", svc.GetName()),
			Path:        path + "/{id}",
			Method:      http.MethodDelete,
		}, svc.DeleteSingle)
	}
	if config.DeleteMode <= BulkSingle {
		slog.Debug("Registering DeleteBulk operation", slog.String("path", path))
		huma.Register(api, huma.Operation{
			OperationID: fmt.Sprintf("delete-bulk-%s", svc.GetName()),
			Summary:     fmt.Sprintf("Delete bulk-%s", svc.GetName()),
			Description: fmt.Sprintf("Batch deletion operation for multiple %s resources. This operation cannot be undone.", svc.GetName()),
			Path:        path,
			Method:      http.MethodDelete,
		}, svc.DeleteBulk)
	}
}

// NewSQLRepository initializes a repository based on the SQL database driver.
func NewSQLRepository[Model any](db *sql.DB) repository.Repository[Model] {
	// Determine the database driver and initialize the appropriate repository
	driverType := reflect.ValueOf(db.Driver()).Type().String()
	slog.Debug("Initializing SQL repository", slog.String("driver", driverType))

	switch driverType {
	case "*mysql.MySQLDriver":
		slog.Debug("Using MySQL repository")
		return repository.NewMySQLRepository[Model](db)
	case "*pq.Driver", "pqx.Driver":
		slog.Debug("Using Postgres repository")
		return repository.NewPostgresRepository[Model](db)
	case "*sqlite3.SQLiteDriver", "*driver.SQLite":
		slog.Debug("Using SQLite repository")
		return repository.NewSQLiteRepository[Model](db)
	case "*mssql.Driver":
		slog.Debug("Using MSSQL repository")
		return repository.NewMSSQLRepository[Model](db)
	}

	slog.Error("Unsupported database driver", slog.String("driver", driverType))
	panic("unsupported database driver " + driverType)
}
