# Getting Started

## Installation

Install GoCRUD using `go get`:

```bash
go get github.com/ckoliber/gocrud
```

## Basic Usage

### 1. Define Your Model

```go
type User struct {
    _    struct{} `db:"users" json:"-"`     // Table name
    ID   *int     `db:"id" json:"id"`       // Primary key
    Name *string  `db:"name" json:"name"`   // Regular field
    Age  *int     `db:"age" json:"age"`     // Regular field
}
```

### 2. Initialize Database

```go
import (
    "database/sql"
    _ "github.com/lib/pq"  // Or any other database driver
)

db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/dbname")
if err != nil {
    panic(err)
}
```

### 3. Register Your API

```go
import (
    "github.com/ckoliber/gocrud"
    "github.com/danielgtaylor/huma/v2"
)

func main() {
    api := huma.New("My API", "1.0.0")

    // Create repository and register routes
    repo := gocrud.NewSQLRepository[User](db)
    gocrud.Register(api, repo, &gocrud.Config[User]{})

    api.Serve()
}
```

## Available Endpoints

Your API now has these endpoints:

-   `GET /users` - List users (with filtering, pagination)
-   `GET /users/{id}` - Get single user
-   `PUT /users` - Update multiple users
-   `PUT /users/{id}` - Update user
-   `POST /users` - Create multiple users
-   `POST /users/one` - Create single user
-   `DELETE /users` - Delete multiple users (with filtering)
-   `DELETE /users/{id}` - Delete user

## Query Parameters

### Filtering

Use the `where` parameter for filtering:

```http
GET /users?where={"age":{"_gt":18}}
GET /users?where={"name":{"_like":"John%"}}
```

### Pagination

Use `limit` and `skip` for pagination:

```http
GET /users?limit=10&skip=20
```

### Sorting

Use `order` for sorting:

```http
GET /users?order={"name":"ASC","age":"DESC"}
```

## Advanced Models

### Relations

Define models with relationships:

```go
type User struct {
    _         struct{}   `db:"users" json:"-"`
    ID        *int       `db:"id" json:"id"`
    Name      *string    `db:"name" json:"name"`
    Documents []Document `db:"documents" src:"id" dest:"userId" table:"documents" json:"-"`
}

type Document struct {
    _      struct{} `db:"documents" json:"-"`
    ID     *int     `db:"id" json:"id"`
    Title  string   `db:"title" json:"title"`
    UserID int      `db:"userId" json:"userId"`
}
```

Query through relations:

```http
GET /users?where={"documents":{"title":{"_eq":"Report"}}}
```

### Custom Operations

Add custom filtering operations:

```go
type ID int

func (_ *ID) Operations() map[string]func(string, ...string) string {
    return map[string]func(string, ...string) string{
        "_regexp": func(key string, values ...string) string {
            return fmt.Sprintf("%s REGEXP %s", key, values[0])
        },
    }
}

// Use in queries
GET /users?where={"id":{"_regexp":"^10.*"}}
```

## Next Steps

-   Check out [CRUD Operations](crud-operations.md) for detailed API usage
-   Learn about [Configuration](configuration.md) options
-   Explore [CRUD Hooks](crud-hooks.md) for custom logic
