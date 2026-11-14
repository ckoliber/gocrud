# Migrations example

gocrud maps your Go structs to an existing schema, it never creates or changes tables. Keep the schema as versioned SQL in [`migrations/`](migrations) and apply it with [golang-migrate](https://github.com/golang-migrate/migrate) before starting the server. This example uses SQLite so it runs without a database server; golang-migrate supports Postgres, MySQL and SQL Server the same way with a different `-database` URL.

## Run

From this directory:

```sh
# Apply the migrations (creates locked.sqlite)
go run -tags sqlite3 github.com/golang-migrate/migrate/v4/cmd/migrate@v4.20.1 \
  -path migrations -database sqlite3://locked.sqlite up

# Start the API on :8888
go run .
```

Try it:

```sh
curl -X POST localhost:8888/user -H 'Content-Type: application/json' -d '[{"name":"David","age":25}]'
curl localhost:8888/user
```

## Migrations

To install the CLI instead of using `go run`:

```sh
go install -tags sqlite3 github.com/golang-migrate/migrate/v4/cmd/migrate@v4.20.1
```

Common commands:

```sh
migrate -path migrations -database sqlite3://locked.sqlite version    # current version
migrate -path migrations -database sqlite3://locked.sqlite down 1     # roll back the last migration
migrate create -ext sql -dir migrations -seq add_users_email          # new up/down pair
```

To change the schema, add a new up/down pair to `migrations/`, run `up`, then update the `User` struct to match.
