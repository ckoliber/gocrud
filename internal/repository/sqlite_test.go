package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQLiteRepository(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INTEGER
		)
	`)
	if err != nil {
		panic(err)
	}

	repo := NewSQLiteRepository[User](db)

	UnitTests(context.Background(), t, repo)
}
