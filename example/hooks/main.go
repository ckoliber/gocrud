package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

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

func main() {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

	api.UseMiddleware()

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	gocrud.Register(api, gocrud.NewSQLRepository[User](db), &gocrud.Config[User]{
		BeforeGet: func(ctx context.Context, where *map[string]any, order *map[string]any, limit *int, skip *int) error {
			if *limit > 50 {
				*limit = 50
			}

			return nil
		},
		BeforePut: func(ctx context.Context, models *[]User) error {
			return nil
		},
		BeforePost: func(ctx context.Context, models *[]User) error {
			return nil
		},
		BeforeDelete: func(ctx context.Context, where *map[string]any) error {
			return nil
		},
		AfterGet: func(ctx context.Context, models *[]User) error {
			return nil
		},
		AfterPut: func(ctx context.Context, models *[]User) error {
			return nil
		},
		AfterPost: func(ctx context.Context, models *[]User) error {
			return nil
		},
		AfterDelete: func(ctx context.Context, models *[]User) error {
			return nil
		},
	})

	fmt.Printf("Starting server on port 8888...\n")
	http.ListenAndServe(":8888", mux)
}
