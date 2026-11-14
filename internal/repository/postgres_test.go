package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestPostgresRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping container test in short mode")
	}
	ctx := context.Background()

	container, err := postgres.Run(ctx, "postgres:18-alpine", postgres.BasicWaitStrategies())
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err)
	defer db.Close()

	createUsersTable(t, db, `CREATE TABLE users (id SERIAL PRIMARY KEY, name TEXT, age INTEGER)`)

	UnitTests(ctx, t, NewPostgresRepository[User](db))
}
