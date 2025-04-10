package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ckoliber/gocrud"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/autopatch"
	"github.com/danielgtaylor/huma/v2/humacli"

	_ "github.com/lib/pq"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

type User struct {
	_         struct{}   `db:"users" json:"-"`
	ID        *int       `db:"id" json:"id" required:"false"`
	Name      *string    `db:"name" json:"name" required:"false" maxLength:"30" example:"David" doc:"User name"`
	Age       *int       `db:"age" json:"age" required:"false" minimum:"1" maximum:"120" example:"25" doc:"User age from 1 to 120"`
	Documents []Document `db:"id,userId" json:"-,documents"`
}

type Document struct {
	_       struct{} `db:"documents" json:"-"`
	ID      *int     `db:"id" json:"id" required:"false"`
	Title   string   `db:"title" json:"title" maxLength:"50" doc:"Document title"`
	Content string   `db:"content" json:"content" maxLength:"500" doc:"Document content"`
	UserID  int      `db:"userId" json:"userId" doc:"Document userId"`
	User    User     `db:"userId,id" json:"-,users"`
}

func main() {
	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		mux := http.NewServeMux()
		api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

		api.UseMiddleware()

		db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=postgres sslmode=disable")
		if err != nil {
			fmt.Println(err)
		}

		gocrud.Register(api, gocrud.NewSQLRepository[User](db), &gocrud.Config[User]{})
		gocrud.Register(api, gocrud.NewSQLRepository[Document](db), &gocrud.Config[Document]{})
		autopatch.AutoPatch(api)

		// Create the HTTP server.
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", options.Port),
			Handler: mux,
		}

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			server.ListenAndServe()
		})

		// Tell the CLI how to stop your server.
		hooks.OnStop(func() {
			fmt.Printf("Stopping server on port %d...\n", options.Port)
			// Give the server 5 seconds to gracefully shut down, then give up.
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
