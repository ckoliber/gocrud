# CRUD Hooks

CRUD hooks allow you to add custom logic that executes before or after CRUD operations. This guide explains how to use hooks effectively in your GoCRUD applications.

## Available Hooks

GoCRUD provides both "before" and "after" hooks for each CRUD operation:

### Before Hooks

-   `BeforeGet`: Executes before retrieving resources
-   `BeforePut`: Executes before updating resources
-   `BeforePost`: Executes before creating resources
-   `BeforeDelete`: Executes before deleting resources

### After Hooks

-   `AfterGet`: Executes after retrieving resources
-   `AfterPut`: Executes after updating resources
-   `AfterPost`: Executes after creating resources
-   `AfterDelete`: Executes after deleting resources

## Hook Signatures

Each hook type has a specific function signature:

```go
// Get operation hooks
BeforeGet func(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) error
AfterGet  func(ctx context.Context, models *[]Model) error

// Put operation hooks
BeforePut func(ctx context.Context, models *[]Model) error
AfterPut  func(ctx context.Context, models *[]Model) error

// Post operation hooks
BeforePost func(ctx context.Context, models *[]Model) error
AfterPost  func(ctx context.Context, models *[]Model) error

// Delete operation hooks
BeforeDelete func(ctx context.Context, where *map[string]any) error
AfterDelete  func(ctx context.Context, models *[]Model) error
```

## Using Hooks

### Basic Hook Configuration

Here's how to configure hooks when registering your model:

```go
gocrud.Register(api, repo, &gocrud.Config[User]{
    BeforePost: func(ctx context.Context, models *[]User) error {
        // Add validation logic here
        return nil
    },
    AfterPost: func(ctx context.Context, models *[]User) error {
        // Add post-processing logic here
        return nil
    },
})
```

### Common Use Cases

#### Input Validation

```go
BeforePost: func(ctx context.Context, models *[]User) error {
    for _, user := range *models {
        if user.Age != nil && *user.Age < 18 {
            return fmt.Errorf("users must be 18 or older")
        }
    }
    return nil
}
```

#### Data Enrichment

```go
AfterGet: func(ctx context.Context, models *[]User) error {
    for i := range *models {
        if (*models)[i].Name == nil {
            defaultName := "Anonymous"
            (*models)[i].Name = &defaultName
        }
    }
    return nil
}
```

#### Access Control

```go
BeforeGet: func(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) error {
    userID := ctx.Value("userID").(string)
    if userID == "" {
        return fmt.Errorf("unauthorized")
    }
    return nil
}
```

#### Audit Logging

```go
AfterDelete: func(ctx context.Context, models *[]User) error {
    for _, user := range *models {
        log.Printf("User deleted: ID=%v", *user.ID)
    }
    return nil
}
```

## Hook Execution Order

1. Before hooks execute first, allowing you to:

    - Validate input
    - Modify query parameters
    - Check permissions
    - Cancel the operation by returning an error

2. The main operation executes only if the before hook succeeds

3. After hooks execute last, allowing you to:
    - Modify returned data
    - Trigger side effects
    - Log operations
    - Send notifications

## Error Handling

-   Any error returned from a hook will stop the operation
-   Before hook errors prevent the main operation from executing
-   After hook errors are returned to the client even though the main operation succeeded

Example error handling:

```go
BeforePut: func(ctx context.Context, models *[]User) error {
    for _, user := range *models {
        if err := validateUser(user); err != nil {
            return fmt.Errorf("validation failed: %w", err)
        }
    }
    return nil
}
```

## Best Practices

1. **Keep Hooks Focused**: Each hook should have a single responsibility

2. **Handle Errors Gracefully**: Always return meaningful error messages

3. **Use Context**: Leverage context for request-scoped data

4. **Consider Performance**: Avoid expensive operations in hooks

5. **Be Careful with Mutations**:
    - Before hooks: Modify input parameters only when necessary
    - After hooks: Be cautious when modifying returned data

## Example: Complete Hook Configuration

```go
config := &gocrud.Config[User]{
    BeforeGet: func(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) error {
        // Add query restrictions
        return nil
    },
    AfterGet: func(ctx context.Context, models *[]User) error {
        // Enrich returned data
        return nil
    },
    BeforePost: func(ctx context.Context, models *[]User) error {
        // Validate new resources
        return nil
    },
    AfterPost: func(ctx context.Context, models *[]User) error {
        // Send notifications
        return nil
    },
    BeforePut: func(ctx context.Context, models *[]User) error {
        // Validate updates
        return nil
    },
    AfterPut: func(ctx context.Context, models *[]User) error {
        // Log changes
        return nil
    },
    BeforeDelete: func(ctx context.Context, where *map[string]any) error {
        // Check delete permissions
        return nil
    },
    AfterDelete: func(ctx context.Context, models *[]User) error {
        // Cleanup related resources
        return nil
    },
}
```
