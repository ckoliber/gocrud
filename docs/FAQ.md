# FAQ

## General Questions

### What databases does GoCRUD support?

GoCRUD supports:

-   PostgreSQL
-   MySQL
-   SQLite
-   Microsoft SQL Server

### Does GoCRUD support PATCH operations?

Yes, through Huma's autopatch feature. Enable it with:

```go
autopatch.AutoPatch(api)
```

### Can I use custom field types?

Yes, by implementing the `Operations` method for custom filtering:

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

## Common Issues

### Why am I getting "unsupported database driver"?

Make sure you:

1. Import the correct database driver
2. Use the supported driver package:
    - PostgreSQL: `github.com/lib/pq`
    - MySQL: `github.com/go-sql-driver/mysql`
    - SQLite: `github.com/mattn/go-sqlite3`
    - MSSQL: `github.com/microsoft/go-mssqldb`

### Why aren't my relations working?

Check that you:

1. Used the correct tag format:
    ```go
    Documents []Document `db:"documents" src:"id" dest:"userId" table:"documents" json:"-"`
    ```
2. Set up the foreign key fields correctly
3. Have proper database permissions

### Why isn't pagination working?

Ensure you're using both `limit` and `skip`:

```http
GET /users?limit=10&skip=0
```

## Best Practices

### How should I structure my models?

Follow these guidelines:

1. Always define table name using the `_` field:
    ```go
    _    struct{} `db:"users" json:"-"`
    ```
2. Make ID fields pointers for proper null handling:
    ```go
    ID   *int    `db:"id" json:"id"`
    ```
3. Use appropriate tags for validation:
    ```go
    Age  *int    `db:"age" json:"age" minimum:"0" maximum:"120"`
    ```

### How can I optimize performance?

1. Use bulk operations when possible
2. Set appropriate limits on queries
3. Add database indexes for filtered fields
4. Use relationship filtering judiciously

### How should I handle errors?

1. Use hooks for validation:
    ```go
    BeforePost: func(ctx context.Context, models *[]User) error {
        if err := validate(models); err != nil {
            return fmt.Errorf("validation failed: %w", err)
        }
        return nil
    }
    ```
2. Return specific error types
3. Log errors appropriately

## Configuration

### How do I disable certain operations?

Use operation modes in config:

```go
config := &gocrud.Config[User]{
    GetMode:    gocrud.BulkSingle,
    PostMode:   gocrud.None,      // Disable POST
    DeleteMode: gocrud.None,      // Disable DELETE
}
```

### How do I add custom validation?

Use before hooks:

```go
BeforePost: func(ctx context.Context, models *[]User) error {
    for _, user := range *models {
        if err := customValidation(user); err != nil {
            return err
        }
    }
    return nil
}
```

## Advanced Usage

### Can I use transactions?

Transactions are handled automatically for:

-   Bulk updates
-   Bulk deletes
-   Operations with hooks

### How do I implement custom filtering?

1. Define custom operations on field types
2. Use the operations in queries:
    ```http
    GET /users?where={"id":{"_regexp":"^10.*"}}
    ```

### Can I extend the default operations?

Yes, by:

1. Implementing custom field types
2. Adding hooks for custom logic
3. Using the underlying repository interface

## Troubleshooting

### Common Error Messages

#### "entity not found"

-   Check if the resource exists
-   Verify the ID is correct
-   Ensure proper database permissions

#### "invalid identifier type"

-   Check model field types
-   Ensure ID fields are properly defined
-   Verify query parameter types

#### "validation failed"

-   Check input data format
-   Verify field constraints
-   Look for missing required fields

### Debug Tips

1. Enable debug logging
2. Check SQL queries in logs
3. Verify database connection
4. Test queries directly in database
5. Check hook execution order

## Getting Help

### Where can I find more examples?

Check the examples directory in the repository:

-   Basic CRUD
-   Relations
-   Custom operations
-   Different databases
-   Hooks implementation

### How do I report issues?

1. Check existing issues on GitHub
2. Provide minimal reproduction
3. Include:
    - Go version
    - Database type and version
    - Complete error message
    - Example code
