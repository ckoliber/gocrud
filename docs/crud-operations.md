# CRUD Operations

This guide details the CRUD operations available in GoCRUD and how to use them effectively.

## GET Operations

### Get Single Resource

Retrieves a single resource by its ID.

```http
GET /users/{id}
```

Response:

```json
{
    "body": {
        "id": 1,
        "name": "John Doe",
        "age": 30
    }
}
```

### Get Multiple Resources

Retrieves multiple resources with filtering, sorting, and pagination.

```http
GET /users?where={"age":{"_gt":25}}&order={"name":"ASC"}&limit=10&skip=0
```

Response:

```json
{
    "body": [
        {
            "id": 1,
            "name": "Alice",
            "age": 28
        },
        {
            "id": 2,
            "name": "Bob",
            "age": 32
        }
    ]
}
```

#### Query Parameters

-   `where`: JSON object for filtering
-   `order`: JSON object for sorting
-   `limit`: Maximum number of items to return
-   `skip`: Number of items to skip

#### Filtering Operators

-   `_eq`: Equal to
-   `_neq`: Not equal to
-   `_gt`: Greater than
-   `_gte`: Greater than or equal to
-   `_lt`: Less than
-   `_lte`: Less than or equal to
-   `_like`: LIKE pattern matching
-   `_nlike`: NOT LIKE pattern matching
-   `_ilike`: Case-insensitive LIKE
-   `_nilike`: Case-insensitive NOT LIKE
-   `_in`: In array
-   `_nin`: Not in array

## POST Operations

### Create Single Resource

Creates a single resource.

```http
POST /users/one
Content-Type: application/json

{
    "body": {
        "name": "John Doe",
        "age": 30
    }
}
```

Response:

```json
{
    "body": {
        "id": 1,
        "name": "John Doe",
        "age": 30
    }
}
```

### Create Multiple Resources

Creates multiple resources in a single request.

```http
POST /users
Content-Type: application/json

{
    "body": [
        {
            "name": "John Doe",
            "age": 30
        },
        {
            "name": "Jane Smith",
            "age": 25
        }
    ]
}
```

Response:

```json
{
    "body": [
        {
            "id": 1,
            "name": "John Doe",
            "age": 30
        },
        {
            "id": 2,
            "name": "Jane Smith",
            "age": 25
        }
    ]
}
```

## PUT Operations

### Update Single Resource

Updates a single resource by its ID.

```http
PUT /users/{id}
Content-Type: application/json

{
    "body": {
        "name": "John Smith",
        "age": 31
    }
}
```

Response:

```json
{
    "body": {
        "id": 1,
        "name": "John Smith",
        "age": 31
    }
}
```

### Update Multiple Resources

Updates multiple resources in a single request.

```http
PUT /users
Content-Type: application/json

{
    "body": [
        {
            "id": 1,
            "name": "John Smith",
            "age": 31
        },
        {
            "id": 2,
            "name": "Jane Doe",
            "age": 26
        }
    ]
}
```

Response:

```json
{
    "body": [
        {
            "id": 1,
            "name": "John Smith",
            "age": 31
        },
        {
            "id": 2,
            "name": "Jane Doe",
            "age": 26
        }
    ]
}
```

## DELETE Operations

### Delete Single Resource

Deletes a single resource by its ID.

```http
DELETE /users/{id}
```

Response:

```json
{
    "body": {
        "id": 1,
        "name": "John Smith",
        "age": 31
    }
}
```

### Delete Multiple Resources

Deletes multiple resources based on filtering criteria.

```http
DELETE /users?where={"age":{"_lt":25}}
```

Response:

```json
{
    "body": [
        {
            "id": 3,
            "name": "Alice Young",
            "age": 22
        },
        {
            "id": 4,
            "name": "Bob Junior",
            "age": 21
        }
    ]
}
```

## Advanced Queries

### Relation Filtering

Filter resources based on related entities:

```http
GET /users?where={"documents":{"title":{"_like":"Report%"}}}
```

### Custom Operations

Use custom field operations if defined:

```http
GET /users?where={"id":{"_regexp":"^10.*"}}
```

### Complex Filters

Combine multiple conditions:

```http
GET /users?where={"_and":[{"age":{"_gt":20}},{"name":{"_like":"J%"}}]}
```

Use OR conditions:

```http
GET /users?where={"_or":[{"age":{"_lt":20}},{"age":{"_gt":60}}]}
```

Use NOT conditions:

```http
GET /users?where={"_not":{"age":{"_eq":30}}}
```

## Error Handling

Common error responses:

-   `400 Bad Request`: Invalid input data
-   `404 Not Found`: Resource not found
-   `422 Unprocessable Entity`: Validation error
-   `500 Internal Server Error`: Server error

Error response format:

```json
{
    "error": {
        "code": "NOT_FOUND",
        "message": "entity not found"
    }
}
```
