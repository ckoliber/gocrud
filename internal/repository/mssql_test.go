package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mssql"
)

func TestMSSQLRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping container test in short mode")
	}
	ctx := context.Background()

	container, err := mssql.Run(ctx, "mcr.microsoft.com/mssql/server:2022-latest", mssql.WithAcceptEULA(), mssql.WithPassword("Gocrud-Test-1"))
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	dsn, err := container.ConnectionString(ctx)
	require.NoError(t, err)
	db, err := sql.Open("sqlserver", dsn)
	require.NoError(t, err)
	defer db.Close()

	createUsersTable(t, db, `CREATE TABLE users (id INT IDENTITY(1,1) PRIMARY KEY, name NVARCHAR(255), age INT)`)

	UnitTests(ctx, t, NewMSSQLRepository[User](db))
}
