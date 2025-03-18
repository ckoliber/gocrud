# GoCRUD

A powerful Go module that extends [Huma](https://huma.rocks/) to automatically generate CRUD APIs with built-in support for permissions and relations.

## 🚀 Features

-   Seamless integration with Huma API framework
-   Automatic CRUD endpoint generation from your models
-   Built-in permissions system
-   Relationship handling (one-to-one, one-to-many, many-to-many)
-   Input validation out of the box
-   Customizable API behaviors
-   Clean and maintainable generated code
-   Compatible with existing Huma middleware

## 📋 Prerequisites

-   Go 1.16 or higher
-   A project using [Huma](https://huma.rocks/)

## 🛠️ Installation

```bash
go get github.com/ckoliber/gocrud
```

## 🎯 Quick Start

1. In your existing Huma project, define your model:

```go
type User struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}
```

2. Register your model with GoCRUD in your main application:

```go
package main

import (
    "github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

    "github.com/ckoliber/gocrud"
)

func main() {
    mux := http.NewServeMux()

    // Initialize your Huma API
    api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

    // Register your model with GoCRUD
    gocrud.Register[User](api, &gocrud.Config{
        // Optional: Configure permissions
        Permissions: &gocrud.Permissions{
            Create: []string{"admin", "user"},
            Read:   []string{"admin", "user"},
            Update: []string{"admin"},
            Delete: []string{"admin"},
        },
        // Optional: Configure relations
        Relations: []gocrud.Relation{
            {
                Field: "Posts",
                Type:  "Post",
                Kind:  gocrud.OneToMany,
            },
        },
    })

    // Start your API server
    http.ListenAndServe(fmt.Sprintf(":%d", options.Port), mux)
}
```

## 📚 Generated Endpoints

GoCRUD automatically generates the following RESTful endpoints for your model:

-   `GET /users` - List all users (with pagination)
-   `POST /users` - Create a new user
-   `GET /users/{id}` - Get a specific user
-   `PUT /users/{id}` - Update a user
-   `DELETE /users/{id}` - Delete a user
-   `GET /users/{id}/posts` - Get related posts (when relations are configured)

## 🔧 Configuration Options

```go
type Config struct {
    // Define who can perform which operations
    Permissions *Permissions

    // Configure model relationships
    Relations []Relation

    // Custom base path (default: plural of model name)
    BasePath string

    // Custom validators
    Validators map[string]ValidatorFunc

    // Hooks for custom logic
    Hooks Hooks
}

type Hooks struct {
    BeforeCreate func(ctx context.Context, model interface{}) error
    AfterCreate  func(ctx context.Context, model interface{}) error
    // ... more hooks available
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

## ✨ Acknowledgments

-   Built on top of the excellent [Huma](https://huma.rocks/) framework
-   Inspired by best practices in the Go community
-   Thanks to all contributors who have helped shape GoCRUD

---

Made with ❤️ by KoLiBer
