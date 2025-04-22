package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/ckoliber/gocrud"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	_ "github.com/lib/pq"
)

type User struct {
	_    struct{} `db:"users" json:"-"`
	ID   *int     `db:"id" json:"id" required:"false"`
	Name *string  `db:"name" json:"name" required:"false" maxLength:"30" example:"David" doc:"User name"`
	Age  *int     `db:"age" json:"age" required:"false" minimum:"1" maximum:"120" example:"25" doc:"User age from 1 to 120"`
}

func NewAuthMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		parts := strings.Split(ctx.Operation().OperationID, "-")

		// Skip auth for GET operations
		if parts[0] == "get" {
			next(ctx)
			return
		}

		// Check for the Authorization header
		token := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
		if len(token) == 0 {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Extract role from the token
		if token == "" {
			huma.WriteErr(api, ctx, http.StatusForbidden, "Forbidden")
			return
		}
		ctx = huma.WithValue(ctx, "user", "12345")
		ctx = huma.WithValue(ctx, "role", "admin")

		next(ctx)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

	api.UseMiddleware(NewAuthMiddleware(api))

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	gocrud.Register(api, gocrud.NewSQLRepository[User](db), &gocrud.Config[User]{
		BeforeDelete: func(ctx context.Context, where *map[string]any) error {
			if ctx.Value("role") == "admin" {
				return nil
			}

			*where = map[string]any{
				"_and": []map[string]any{
					{"id": map[string]any{"_eq": ctx.Value("user")}},
					*where,
				},
			}

			return nil
		},
	})

	fmt.Printf("Starting server on port 8888...\n")
	http.ListenAndServe(":8888", mux)
}
