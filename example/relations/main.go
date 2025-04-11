package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ckoliber/gocrud"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	_ "github.com/lib/pq"
)

type User struct {
	_         struct{}   `db:"users" json:"-"`
	ID        *int       `db:"id" json:"id" required:"false"`
	Name      *string    `db:"name" json:"name" required:"false" maxLength:"30" example:"David" doc:"User name"`
	Age       *int       `db:"age" json:"age" required:"false" minimum:"1" maximum:"120" example:"25" doc:"User age from 1 to 120"`
	Documents []Document `db:"documents" src:"id" dest:"userId" table:"documents" json:"-"`
}

type Document struct {
	_       struct{} `db:"documents" json:"-"`
	ID      *int     `db:"id" json:"id" required:"false"`
	Title   string   `db:"title" json:"title" maxLength:"50" doc:"Document title"`
	Content string   `db:"content" json:"content" maxLength:"500" doc:"Document content"`
	UserID  int      `db:"userId" json:"userId" doc:"Document userId"`
	User    User     `db:"user" src:"userId" dest:"id" table:"users" json:"-"`
}

func main() {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

	api.UseMiddleware()

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	gocrud.Register(api, gocrud.NewSQLRepository[User](db), &gocrud.Config[User]{})
	gocrud.Register(api, gocrud.NewSQLRepository[Document](db), &gocrud.Config[Document]{})

	fmt.Printf("Starting server on port 8888...\n")
	http.ListenAndServe(":8888", mux)
}
