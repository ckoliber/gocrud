package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSQLiteRepository(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:repository?mode=memory&cache=shared")
	require.NoError(t, err)
	defer db.Close()

	createUsersTable(t, db, `CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER)`)

	UnitTests(context.Background(), t, NewSQLiteRepository[User](db))
}
