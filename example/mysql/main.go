package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ckoliber/gocrud"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	_ "github.com/go-sql-driver/mysql"
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

	db, err := sql.Open("mysql", "root:password@/default")
	if err != nil {
		fmt.Println(err)
	}

	gocrud.Register(api, gocrud.NewSQLRepository[User](db), &gocrud.Config[User]{})

	fmt.Printf("Starting server on port 8888...\n")
	http.ListenAndServe(":8888", mux)
}
