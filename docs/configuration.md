# Configuration

This guide explains how to configure GoCRUD for your application's needs.

## Core Configuration

When registering your models with GoCRUD, you can provide configuration options through the `Config` struct:

```go
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
```

## Operation Modes

GoCRUD supports three operation modes for each CRUD operation:

```go
type Mode int

const (
    BulkSingle Mode = iota  // Both bulk and single operations enabled
    Single                  // Only single operations enabled
    None                    // Operation disabled
)
```

Example configuration:

```go
config := &gocrud.Config[User]{
    GetMode:    gocrud.BulkSingle,  // Enable both GET /users and GET /users/{id}
    PutMode:    gocrud.Single,      // Enable only PUT /users/{id}
    PostMode:   gocrud.BulkSingle,  // Enable both POST /users and POST /users/one
    DeleteMode: gocrud.None,        // Disable all DELETE operations
}
```

## Hook Configuration

Hooks allow you to add custom logic before and after CRUD operations.

### Before Hooks

Before hooks run before the database operation:

```go
config := &gocrud.Config[User]{
    BeforePost: func(ctx context.Context, models *[]User) error {
        // Validate age before creating users
        for _, user := range *models {
            if user.Age != nil && *user.Age < 18 {
                return fmt.Errorf("users must be 18 or older")
            }
        }
        return nil
    },
}
```

### After Hooks

After hooks run after the database operation but before the response is sent:

```go
config := &gocrud.Config[User]{
    AfterGet: func(ctx context.Context, models *[]User) error {
        // Modify or enrich user data
        for i := range *models {
            if (*models)[i].Name == nil {
                defaultName := "Anonymous"
                (*models)[i].Name = &defaultName
            }
        }
        return nil
    },
}
```

## Model Configuration

Models are configured using struct tags:

```go
type User struct {
    _    struct{} `db:"users" json:"-"`                    // Table name
    ID   *int     `db:"id" json:"id" required:"false"`     // Primary key field
    Name *string  `db:"name" json:"name,omitempty"`        // Regular field
    Age  *int     `db:"age" json:"age" minimum:"0"`        // Field with validation
}
```

### Available Tags

GoCRUD uses the following core tags:

-   `db`: Database column name or table name (on the `_` field)
-   `json`: JSON field name and options (e.g., `-` or `omitempty`)
-   `src`: Source field name in relationships
-   `dest`: Destination field name in relationships
-   `table`: Related table name in relationships

Additional validation tags (like `required`, `minimum`, `maximum`, etc.) are available through the [Huma framework validation tags](https://huma.rocks/).

### Relation Configuration

For related models, additional tags are used:

```go
type User struct {
    _         struct{}   `db:"users" json:"-"`
    ID        *int       `db:"id" json:"id"`
    Documents []Document `db:"documents" src:"id" dest:"userId" table:"documents" json:"-"`
}
```

Relation tags:

-   `src`: Source field name in the current table
-   `dest`: Destination field name in the related table
-   `table`: Related table name

## Custom Field Operations

Define custom operations for field types:

```go
type CustomID int

func (_ *CustomID) Operations() map[string]func(string, ...string) string {
    return map[string]func(string, ...string) string{
        "_regexp": func(key string, values ...string) string {
            return fmt.Sprintf("%s REGEXP %s", key, values[0])
        },
    }
}
```

## Database Configuration

GoCRUD automatically configures itself based on the database driver:

```go
import (
    "database/sql"
    _ "github.com/lib/pq"        // PostgreSQL
    _ "github.com/go-sql-driver/mysql"  // MySQL
    _ "github.com/mattn/go-sqlite3"     // SQLite
    _ "github.com/microsoft/go-mssqldb" // MSSQL
)

// Database connection
db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/dbname")

// Repository creation
repo := gocrud.NewSQLRepository[User](db)
```

The SQL dialect and parameter style are automatically configured based on the driver.
