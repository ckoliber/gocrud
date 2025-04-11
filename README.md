# GoCRUD

![GoCRUD](./docs/icon.svg)

GoCRUD is a powerful Go module that extends [Huma](https://huma.rocks/) to automatically generate CRUD APIs with built-in support for input validation and customizable hooks. It simplifies API development by automating repetitive tasks, allowing you to focus on your business logic.

## 🚀 Features

-   **Seamless Huma Integration**: Works effortlessly with the Huma API framework.
-   **Automatic CRUD Generation**: Instantly generate RESTful endpoints for your models.
-   **Input Validation**: Automatically validates input data.
-   **Customizable Hooks**: Add custom logic before or after CRUD operations.
-   **Clean and Maintainable**: Keeps your codebase organized and easy to maintain.

## 📋 Prerequisites

-   Go 1.22 or higher
-   A project using [Huma](https://huma.rocks/)

## 🛠️ Installation

Install GoCRUD using `go get`:

```bash
go get github.com/ckoliber/gocrud
```

## 🎯 Quick Start

1. **Define Your Model**:

```go
type User struct {
	_    struct{} `db:"users" json:"-"`
	ID   *int     `db:"id" json:"id" required:"false"`
	Name *string  `db:"name" json:"name" required:"false" maxLength:"30" example:"David" doc:"User name"`
	Age  *int     `db:"age" json:"age" required:"false" minimum:"1" maximum:"120" example:"25" doc:"User age from 1 to 120"`
}
```

2. **Register Your Model with GoCRUD**:

```go
package main

import (
    "github.com/danielgtaylor/huma/v2"
    "github.com/ckoliber/gocrud"
    "database/sql"
    _ "github.com/lib/pq" // Example: PostgreSQL driver
)

func main() {
    db, _ := sql.Open("postgres", "your-dsn-here")
    api := huma.New("My API", "1.0.0")

    repo := gocrud.NewSQLRepository[User](db)
    gocrud.Register(api, repo, &gocrud.Config[User]{})

    api.Serve()
}
```

3. **Run Your API**:

Start your application, and GoCRUD will generate the following endpoints for the `User` model:

-   `GET /users` - List all users
-   `POST /users` - Create a new user
-   `GET /users/{id}` - Get a specific user
-   `PUT /users/{id}` - Update a user
-   `DELETE /users/{id}` - Delete a user

## 🔧 Configuration Options

GoCRUD provides a flexible configuration system to customize API behavior:

```go
type Config[Model any] struct {
    GetMode    Mode // Configure GET behavior (e.g., single or bulk)
    PutMode    Mode // Configure PUT behavior
    PostMode   Mode // Configure POST behavior
    DeleteMode Mode // Configure DELETE behavior

    CRUDHooks[Model] // Add hooks for custom logic
}
```

### Example: Adding Hooks

```go
config := &gocrud.Config[User]{
    CRUDHooks: gocrud.CRUDHooks[User]{
        BeforePost: func(ctx context.Context, models *[]User) error {
            for _, user := range *models {
                if user.Age < 18 {
                    return fmt.Errorf("user must be at least 18 years old")
                }
            }
            return nil
        },
    },
}
```

## 🔰 Advanced Features

### Relation Filtering

GoCRUD supports filtering through relationships. You can query parent entities based on their related entities' properties:

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

You can then filter users by their documents:

```http
GET /users?where={"documents":{"title":{"_eq":"Doc4"}}}
```

This will return users who have documents with title "Doc4".

### Custom Field Operations

You can define custom operations for your model fields by implementing the `Operations` method:

```go
type ID int

func (_ *ID) Operations() map[string]func(string, ...string) string {
    return map[string]func(string, ...string) string{
        "_regexp": func(key string, values ...string) string {
            return fmt.Sprintf("%s REGEXP %s", key, values[0])
        },
        "_iregexp": func(key string, values ...string) string {
            return fmt.Sprintf("%s IREGEXP %s", key, values[0])
        },
    }
}

type User struct {
    _    struct{} `db:"users" json:"-"`
    ID   *ID      `db:"id" json:"id"`
    Name *string  `db:"name" json:"name"`
}
```

Now you can use these custom operations in your queries:

```http
GET /users?where={"id":{"_regexp":"5"}}
```

The operations are type-safe and validated against the field's defined operations.

## 🤝 Contributing

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature/my-feature`.
3. Commit your changes: `git commit -m "Add my feature"`.
4. Push to the branch: `git push origin feature/my-feature`.
5. Open a pull request.

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

## ✨ Acknowledgments

-   Built on top of the [Huma](https://huma.rocks/) framework.
-   Inspired by best practices in the Go community.
-   Thanks to all contributors who have helped shape GoCRUD.

---

Made with ❤️ by KoLiBer
