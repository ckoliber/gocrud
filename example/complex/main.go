package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ckoliber/gocrud"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

type User struct {
	_    struct{} `json:"user"`
	ID   string   `json:"id"`
	Name string   `json:"name" maxLength:"30" example:"David" doc:"User name"`
	Age  int      `json:"age" minimum:"1" maximum:"120" example:"25" doc:"User age from 1 to 120"`
}

type Document struct {
	_       struct{} `json:"document"`
	ID      string   `json:"id"`
	Title   string   `json:"title" maxLength:"50" doc:"Document title"`
	Content string   `json:"content" maxLength:"500" doc:"Document content"`
	UserID  string   `json:"userId" doc:"Document userId"`
}

func main() {
	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		mux := http.NewServeMux()
		api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

		api.UseMiddleware()

		// Register GET /greeting/{name}
		gocrud.Register[User](api)
		gocrud.Register[Document](api)

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
