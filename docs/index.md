# Introduction

![GoCRUD](./icon.svg)

Welcome to the **GoCRUD** documentation! This guide will help you understand how to use GoCRUD to build powerful, scalable, and maintainable CRUD APIs with ease.

## 📖 Overview

GoCRUD is a Go module that extends the [Huma](https://huma.rocks/) framework to provide automatic CRUD API generation. It simplifies API development by automating repetitive tasks, allowing you to focus on your application's business logic.

### Key Features

-   **Automatic CRUD Generation**: Generate RESTful endpoints for your models with minimal configuration.
-   **Input Validation**: Built-in validation for your model fields.
-   **Customizable Hooks**: Add custom logic before or after CRUD operations.
-   **Relationship Filtering**: Query through model relationships with type-safe filters.
-   **Custom Field Operations**: Define custom field-specific filtering operations.
-   **Database Agnostic**: Supports PostgreSQL, MySQL, SQLite, and MSSQL.

### Relations Support

Define relationships between your models and query through them:

```go
type User struct {
    ID        *int       `db:"id" json:"id"`
    Documents []Document `db:"documents" src:"id" dest:"userId" table:"documents" json:"-"`
}

// Query users with specific documents
GET /users?where={"documents":{"title":{"_eq":"Doc4"}}}
```

### Custom Operations

Add custom filtering operations to your field types:

```go
type ID int

func (_ *ID) Operations() map[string]func(string, ...string) string {
    return map[string]func(string, ...string) string{
        "_regexp": func(key string, values ...string) string {
            return fmt.Sprintf("%s REGEXP %s", key, values[0])
        },
    }
}

// Use custom operations in queries
GET /users?where={"id":{"_regexp":"5"}}
```

## 📚 Documentation Structure

The documentation is organized as follows:

-   **[Introduction](introduction.md)**: Learn about GoCRUD and its core concepts.
-   **[Getting Started](getting-started.md)**: Step-by-step guide to set up and use GoCRUD in your project.
-   **[Configuration](configuration.md)**: Detailed explanation of configuration options and examples.
-   **[CRUD Operations](crud-operations.md)**: Learn how to use the generated CRUD endpoints.
-   **[CRUD Hooks](crud-hooks.md)**: Customize your API behavior with hooks.
-   **[Advanced Topics](advanced-topics.md)**: Explore advanced features and future enhancements.
-   **[FAQ](FAQ.md)**: Frequently asked questions and troubleshooting tips.

## 🚀 Getting Started

To get started with GoCRUD, check out the [Getting Started](getting-started.md) guide. It will walk you through the installation process, model definition, and API registration.

## 🛠️ Contributing

We welcome contributions to improve GoCRUD and its documentation. If you'd like to contribute, please check out the [Contributing Guide](https://github.com/ckoliber/gocrud/blob/main/CONTRIBUTING.md).

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/ckoliber/gocrud/blob/main/LICENSE.md) file for details.

---

Made with ❤️ by KoLiBer
