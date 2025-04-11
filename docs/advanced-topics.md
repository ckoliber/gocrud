# Advanced Topics

This guide covers advanced features and usage patterns in GoCRUD.

## Model Relations

GoCRUD supports one-to-one and one-to-many relationships between models.

### Defining Relations

```go
type User struct {
    _         struct{}   `db:"users" json:"-"`
    ID        *int       `db:"id" json:"id"`
    Documents []Document `db:"documents" src:"id" dest:"userId" table:"documents" json:"-"`  // One-to-many
}

type Document struct {
    _      struct{} `db:"documents" json:"-"`
    ID     *int     `db:"id" json:"id"`
    UserID int      `db:"userId" json:"userId"`
    User   User     `db:"user" src:"userId" dest:"id" table:"users" json:"-"`    // One-to-one
}
```

Relation tags:

-   `db`: Name of the related table
-   `src`: Source field in the current model
-   `dest`: Destination field in the related model
-   `table`: Target table name
-   `json`: Usually "-" to exclude from JSON

### Querying Relations

Filter records based on related entities:

```http
# Find users who have documents with specific titles
GET /users?where={"documents":{"title":{"_eq":"Report"}}}

# Find documents belonging to users of a certain age
GET /documents?where={"user":{"age":{"_gt":30}}}
```

## Custom Field Operations

Define custom filtering operations for specific field types:

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
    ID   *ID     `db:"id" json:"id"`
    Name *string `db:"name" json:"name"`
}
```

Using custom operations:

```http
GET /users?where={"id":{"_regexp":"^10.*"}}
```

## Complex Queries

### Logical Operators

Combine multiple conditions:

```http
# AND operator
GET /users?where={"_and":[{"age":{"_gt":20}},{"name":{"_like":"J%"}}]}

# OR operator
GET /users?where={"_or":[{"age":{"_lt":20}},{"age":{"_gt":60}}]}

# NOT operator
GET /users?where={"_not":{"age":{"_eq":30}}}
```

### Nested Conditions

Create complex nested queries:

```http
GET /users?where={
    "_or": [
        {
            "_and": [
                {"age": {"_gt": 20}},
                {"name": {"_like": "J%"}}
            ]
        },
        {
            "documents": {
                "title": {"_like": "Report%"}
            }
        }
    ]
}
```

## Performance Optimization

### Query Optimization

1. **Use Appropriate Indexes**:

    - Add indexes for frequently filtered fields
    - Add composite indexes for common filter combinations

2. **Limit Result Sets**:
    - Always use pagination
    - Set reasonable default limits

```http
GET /users?limit=50&skip=0
```

3. **Select Specific Fields**:
    - Coming soon: Field selection support
    - Will allow retrieving only needed fields

### Bulk Operations

Use bulk endpoints for better performance:

```http
# Bulk create
POST /users
{
    "body": [
        {"name": "User1"},
        {"name": "User2"}
    ]
}

# Bulk update
PUT /users
{
    "body": [
        {"id": 1, "name": "Updated1"},
        {"id": 2, "name": "Updated2"}
    ]
}

# Bulk delete
DELETE /users?where={"age":{"_lt":18}}
```
