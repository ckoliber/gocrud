package main

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"net/http"

	"github.com/ckoliber/gocrud"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*.sql
var migrations embed.FS

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

	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		fmt.Println(err)
	}

	// Schema lives in versioned SQL migrations, not in the struct: apply them with golang-migrate
	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		panic(err)
	}
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithInstance("iofs", source, "sqlite3", driver)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	gocrud.Register(api, gocrud.NewSQLRepository[User](db), &gocrud.Config[User]{})

	fmt.Printf("Starting server on port 8888...\n")
	http.ListenAndServe(":8888", mux)
}
